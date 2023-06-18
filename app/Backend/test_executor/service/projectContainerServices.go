package service

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"os"
	"path/filepath"
	

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	natting "github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

func BuildImageforDockerFile2(client *client.Client, tags []string, dockerfile string, contextPath string) error {
	ctx := context.Background()

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	cwd := contextPath

	err := filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(info, path)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, cwd+string(filepath.Separator))

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			if _, err := io.Copy(tw, file); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	dockerFileTarReader := bytes.NewReader(buf.Bytes())

	buildOptions := types.ImageBuildOptions{
		Context:    dockerFileTarReader,
		Dockerfile: dockerfile,
		Remove:     true,
		Tags:       tags,
	}

	imageBuildResponse, err := client.ImageBuild(
		ctx,
		dockerFileTarReader,
		buildOptions,
	)

	if err != nil {
		return err
	}

	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		return err
	}

	return nil

}

func RunContainerForProjectTestImage(client *client.Client, imagename string, containername string, port string, inputEnv []string) error {

	newport, err := natting.NewPort("tcp", port)
	if err != nil {
		fmt.Println("Unable to create docker port")
		return err
	}
	platform := &v1.Platform{}


	hostConfig := &container.HostConfig{
		Links: []string{"Selenium_Container:alias"},
		PortBindings: natting.PortMap{
			newport: []natting.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: port,
				},
			},
		},
		RestartPolicy: container.RestartPolicy{
			Name: "no",
		},
		LogConfig: container.LogConfig{
			Type:   "json-file",
			Config: map[string]string{},
		},
	}

	networkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{},
	}
	gatewayConfig := &network.EndpointSettings{
		Gateway: "gatewayname",
	}
	networkConfig.EndpointsConfig["bridge"] = gatewayConfig

	exposedPorts := map[natting.Port]struct{}{
		newport: {},
	}

	config := &container.Config{
		Image:        imagename,
		Env:          inputEnv,
		ExposedPorts: exposedPorts,
		Hostname:     fmt.Sprintf("%s-hostnameexample", imagename),
	}

	cont, err := client.ContainerCreate(
		context.Background(),
		config,
		hostConfig,
		networkConfig,
		platform,
		containername,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	client.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})
	log.Printf("Container %s is created", cont.ID)

	return nil
}

func StopProjectContainer(client *client.Client, containername string) error {
	ctx := context.Background()

	if err := client.ContainerStop(ctx, containername, nil); err != nil {
		log.Printf("Unable to stop container %s: %s", containername, err)
		return err
	}
	return nil
}

func RemoveProjectContainer(client *client.Client, containername string) error {
	ctx := context.Background()

	removeOptions := types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	}

	if err := client.ContainerRemove(ctx, containername, removeOptions); err != nil {
		log.Printf("Unable to remove container: %s", err)
		return err
	}

	return nil
}
