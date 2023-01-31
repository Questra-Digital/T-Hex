package main

import (
	// "context"
	// "fmt"
	// "io"
	// "log"
	// "os"

	// "github.com/docker/docker/api/types"
	// "github.com/docker/docker/api/types/container"
	// "github.com/docker/docker/api/types/network"
	// "github.com/docker/docker/client"

	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	network "github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	natting "github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

func pullImageSelenium(client *client.Client, image string) error {
	ctx := context.Background()
	reader, err := client.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, reader)
	return nil
}

func runContainerSelenium(client *client.Client, imagename string, containername string, port string, inputEnv []string) error {

	newport, err := natting.NewPort("tcp", port)
	if err != nil {
		fmt.Println("Unable to create docker port")
		return err
	}
	platform := &v1.Platform{}

	hostConfig := &container.HostConfig{
		PortBindings: natting.PortMap{
			newport: []natting.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: port,
				},
			},
		},
		RestartPolicy: container.RestartPolicy{
			Name: "always",
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
		newport: struct{}{},
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

func runContainerForProjectTestImage(client *client.Client, imagename string, containername string, port string, inputEnv []string) error {

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
		newport: struct{}{},
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

func buildImageforDockerFile2(client *client.Client, tags []string, dockerfile string, contextPath string) error {
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

func stopAndRemoveProjectContainer(client *client.Client, containername string) error {
	ctx := context.Background()

	if err := client.ContainerStop(ctx, containername, nil); err != nil {
		log.Printf("Unable to stop container %s: %s", containername, err)
	}

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

func StopsSeleniumContainer(client *client.Client, containername string) error {
	ctx := context.Background()

	if err := client.ContainerStop(ctx, containername, nil); err != nil {
		log.Printf("Unable to stop container %s: %s", containername, err)
	}

	return nil
}

func main() {

	fmt.Println("1) -------Pulling selenium image from docker hub registory-------")

	client, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("Unable to create docker client: %s", err)
	}

	image := "selenium/standalone-chrome"
	err = pullImageSelenium(client, image)
	if err != nil {

		log.Println(err)
		return
	}

	fmt.Println("2) -------Running Selenium_Container-------")

	containername := "Selenium_Container"
	portopening := "4444"
	inputEnv1 := []string{fmt.Sprintf("LISTENINGPORT=%s", portopening)}
	err = runContainerSelenium(client, image, containername, portopening, inputEnv1)
	if err != nil {
		log.Println(err)
	}

	time.Sleep(10 * time.Second)

	fmt.Println("3) -------Create PythonTest Image from a Docker File-------")

	tags1 := []string{"python_test"}
	dockerfile1 := "Dockerfile"
	// contextPath1 := "C:\\Users\\abdul\\Desktop\\T-Hex\\t-hex\\app\\backend\\PythonTest"
	contextPath1 := "example\\PythonTest"
	err = buildImageforDockerFile2(client, tags1, dockerfile1, contextPath1)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println("4) -------Create GoTest Image from a Docker File-------")

	tags2 := []string{"go_test"}
	dockerfile2 := "Dockerfile"
	//contextPath2 := "C:\\Users\\abdul\\Desktop\\T-Hex\\t-hex\\app\\backend\\GoTest"
	contextPath2 := "example\\GoTest"
	err = buildImageforDockerFile2(client, tags2, dockerfile2, contextPath2)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	time.Sleep(10 * time.Second)

	fmt.Println("5) -------Running PythonTest-------")

	containernameproject_test1 := "PythonTest_Container"
	portopening2 := "8080"
	image2 := "python_test"
	inputEnv2 := []string{fmt.Sprintf("LISTENINGPORT=%s", portopening2)}
	err = runContainerForProjectTestImage(client, image2, containernameproject_test1, portopening2, inputEnv2)
	if err != nil {
		log.Println(err)
	}

	time.Sleep(10 * time.Second)

	fmt.Println("6) -------Running GoTest-------")

	containernameproject_test2 := "GoTest_Container"
	portopening3 := "8080"
	image3 := "go_test"
	inputEnv3 := []string{fmt.Sprintf("LISTENINGPORT=%s", portopening3)}
	err = runContainerForProjectTestImage(client, image3, containernameproject_test2, portopening3, inputEnv3)
	if err != nil {
		log.Println(err)
	}

	time.Sleep(40 * time.Second)

	fmt.Println("7) ------- ~(Stop and removing)~ PythonTest_Container-------")

	stopAndRemoveProjectContainer(client, "PythonTest_Container")

	time.Sleep(20 * time.Second)

	fmt.Println("8) ------- ~(Stop and removing)~ GoTest_Container-------")

	stopAndRemoveProjectContainer(client, "GoTest_Container")

	time.Sleep(10 * time.Second)

	fmt.Println("9) -------Stop Selenium_Container-------")
	StopsSeleniumContainer(client, "Selenium_Container")

}
