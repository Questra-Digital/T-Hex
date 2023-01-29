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
	"io/ioutil"
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

func buildImage(client *client.Client, image string) error {
	ctx := context.Background()
	reader, err := client.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, reader)
	return nil
}

func runContainer(client *client.Client, imagename string, containername string, port string, inputEnv []string) error {
	// Define a PORT opening
	newport, err := natting.NewPort("tcp", port)
	if err != nil {
		fmt.Println("Unable to create docker port")
		return err
	}
	platform := &v1.Platform{}
	// Configured hostConfig:
	// https://godoc.org/github.com/docker/docker/api/types/container#HostConfig
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

	// Define Network config (why isn't PORT in here...?:
	// https://godoc.org/github.com/docker/docker/api/types/network#NetworkingConfig
	networkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{},
	}
	gatewayConfig := &network.EndpointSettings{
		Gateway: "gatewayname",
	}
	networkConfig.EndpointsConfig["bridge"] = gatewayConfig

	// Define ports to be exposed (has to be same as hostconfig.portbindings.newport)
	exposedPorts := map[natting.Port]struct{}{
		newport: struct{}{},
	}

	// Configuration
	// https://godoc.org/github.com/docker/docker/api/types/container#Config
	config := &container.Config{
		Image:        imagename,
		Env:          inputEnv,
		ExposedPorts: exposedPorts,
		Hostname:     fmt.Sprintf("%s-hostnameexample", imagename),
	}

	// Creating the actual container. This is "nil,nil,nil" in every example.
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

	// Run the actual container
	client.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})
	log.Printf("Container %s is created", cont.ID)

	return nil
}

func buildImageforDockerFile(client *client.Client, tags []string, dockerfile string) error {
	ctx := context.Background()

	// Create a buffer
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	// Create a filereader
	dockerFileReader, err := os.Open(dockerfile)
	if err != nil {
		return err
	}

	// Read the actual Dockerfile
	readDockerFile, err := ioutil.ReadAll(dockerFileReader)
	if err != nil {
		return err
	}

	// Make a TAR header for the file
	tarHeader := &tar.Header{
		Name: dockerfile,
		Size: int64(len(readDockerFile)),
	}

	// Writes the header described for the TAR file
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		return err
	}

	// Writes the dockerfile data to the TAR file
	_, err = tw.Write(readDockerFile)
	if err != nil {
		return err
	}

	dockerFileTarReader := bytes.NewReader(buf.Bytes())

	// Define the build options to use for the file
	// https://godoc.org/github.com/docker/docker/api/types#ImageBuildOptions
	buildOptions := types.ImageBuildOptions{
		Context:    dockerFileTarReader,
		Dockerfile: dockerfile,
		Remove:     true,
		Tags:       tags,
	}

	// Build the actual image
	imageBuildResponse, err := client.ImageBuild(
		ctx,
		dockerFileTarReader,
		buildOptions,
	)

	if err != nil {
		return err
	}

	// Read the STDOUT from the build process
	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		return err
	}

	return nil
}

func runContainerForProjectTestImage(client *client.Client, imagename string, containername string, port string, inputEnv []string) error {
	// Define a PORT opening
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

	// Define Network config (why isn't PORT in here...?:
	// https://godoc.org/github.com/docker/docker/api/types/network#NetworkingConfig
	networkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{},
	}
	gatewayConfig := &network.EndpointSettings{
		Gateway: "gatewayname",
	}
	networkConfig.EndpointsConfig["bridge"] = gatewayConfig

	// Define ports to be exposed (has to be same as hostconfig.portbindings.newport)
	exposedPorts := map[natting.Port]struct{}{
		newport: struct{}{},
	}

	// Configuration
	// https://godoc.org/github.com/docker/docker/api/types/container#Config
	config := &container.Config{
		Image:        imagename,
		Env:          inputEnv,
		ExposedPorts: exposedPorts,
		Hostname:     fmt.Sprintf("%s-hostnameexample", imagename),
	}

	// Creating the actual container. This is "nil,nil,nil" in every example.
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

	// Run the actual container
	client.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})
	log.Printf("Container %s is created", cont.ID)

	return nil
}

func buildImageforDockerFile1(client *client.Client, tags []string, dockerfile string) error {
	ctx := context.Background()

	// Create a buffer to hold the tar archive
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	// Get the current working directory
	// cwd, err := os.Getwd()
	// if err != nil {
	// 	return err
	// }

	cwd := "C:\\Users\\abdul\\Desktop\\T-Hex\\t-hex\\app\\backend\\PythonTest"

	// Walk the directory and add all files to the tar archive
	err := filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create a tar header for the file
		header, err := tar.FileInfoHeader(info, path)
		if err != nil {
			return err
		}

		// Update the name of the file in the header to be the path relative to the current working directory
		header.Name = strings.TrimPrefix(path, cwd+string(filepath.Separator))

		// Write the header to the tar archive
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// If the file is not a directory, add its contents to the archive
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			// Copy the file's contents to the tar archive
			if _, err := io.Copy(tw, file); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	// Create a reader for the tar archive
	dockerFileTarReader := bytes.NewReader(buf.Bytes())

	// Define the build options to use for the file
	buildOptions := types.ImageBuildOptions{
		Context:    dockerFileTarReader,
		Dockerfile: dockerfile,
		Remove:     true,
		Tags:       tags,
	}

	// Build the actual image
	imageBuildResponse, err := client.ImageBuild(
		ctx,
		dockerFileTarReader,
		buildOptions,
	)

	if err != nil {
		return err
	}

	// Read the STDOUT from the build process
	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		return err
	}

	return nil
}

func buildImageforDockerFile2(client *client.Client, tags []string, dockerfile string) error {
	ctx := context.Background()

	// Create a buffer to hold the tar archive
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	// Get the current working directory
	// cwd, err := os.Getwd()
	// if err != nil {
	// 	return err
	// }

	cwd := "C:\\Users\\abdul\\Desktop\\T-Hex\\t-hex\\app\\backend\\GoTest"

	// Walk the directory and add all files to the tar archive
	err := filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create a tar header for the file
		header, err := tar.FileInfoHeader(info, path)
		if err != nil {
			return err
		}

		// Update the name of the file in the header to be the path relative to the current working directory
		header.Name = strings.TrimPrefix(path, cwd+string(filepath.Separator))

		// Write the header to the tar archive
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// If the file is not a directory, add its contents to the archive
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			// Copy the file's contents to the tar archive
			if _, err := io.Copy(tw, file); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	// Create a reader for the tar archive
	dockerFileTarReader := bytes.NewReader(buf.Bytes())

	// Define the build options to use for the file
	buildOptions := types.ImageBuildOptions{
		Context:    dockerFileTarReader,
		Dockerfile: dockerfile,
		Remove:     true,
		Tags:       tags,
	}

	// Build the actual image
	imageBuildResponse, err := client.ImageBuild(
		ctx,
		dockerFileTarReader,
		buildOptions,
	)

	if err != nil {
		return err
	}

	// Read the STDOUT from the build process
	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		return err
	}

	return nil
}

// Stop and remove a container
func stopAndRemoveContainer(client *client.Client, containername string) error {
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

// Stop and remove a container
func StopsSeleniumContainer(client *client.Client, containername string) error {
	ctx := context.Background()

	if err := client.ContainerStop(ctx, containername, nil); err != nil {
		log.Printf("Unable to stop container %s: %s", containername, err)
	}

	return nil
}

func main() {

	to_check_docker_image := 5

	if to_check_docker_image == 1 {

		client, err := client.NewEnvClient()
		if err != nil {
			log.Fatalf("Unable to create docker client: %s", err)
		}
		fmt.Println("1) -------Pulling selenium image from docker hub registory-------")
		image := "selenium/standalone-chrome"
		err = buildImage(client, image)
		if err != nil {

			log.Println(err)
			return
		}
		fmt.Println("2) -------Running Selenium_Container-------")
		containername := "Selenium_Container"
		portopening := "4444"
		inputEnv1 := []string{fmt.Sprintf("LISTENINGPORT=%s", portopening)}
		err = runContainer(client, image, containername, portopening, inputEnv1)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("3) -------Running ProjectTestContainer-------")
		time.Sleep(30 * time.Second)
		containernameproject_test := "ProjectTestContainer"
		portopening2 := "8080"
		image2 := "my-go-app"
		inputEnv := []string{fmt.Sprintf("LISTENINGPORT=%s", portopening2)}
		err = runContainerForProjectTestImage(client, image2, containernameproject_test, portopening2, inputEnv)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("4) -------Stop and removing ProjectTest Container-------")
		time.Sleep(120 * time.Second)
		stopAndRemoveContainer(client, "ProjectTestContainer")
		fmt.Println("5) -------Stop Selenium_Container Container-------")
		StopsSeleniumContainer(client, "Selenium_Container")

	} else if to_check_docker_image == 2 {

		client, err := client.NewEnvClient()

		fmt.Println("7) -------Running ProjectTestContainer-------")
		time.Sleep(30 * time.Second)
		containernameproject_test := "PythonTestContainer"
		portopening2 := "8080"
		image2 := "command_image"
		inputEnv := []string{fmt.Sprintf("LISTENINGPORT=%s", portopening2)}
		err = runContainerForProjectTestImage(client, image2, containernameproject_test, portopening2, inputEnv)
		if err != nil {
			log.Println(err)
		}

	} else if to_check_docker_image == 3 {

		client, err := client.NewEnvClient()

		fmt.Println("8) -------Running ProjectTestContainer-------")
		time.Sleep(30 * time.Second)
		containernameproject_test := "Go_Lang_Test_Container"
		portopening2 := "8080"
		image2 := "go_check"
		inputEnv := []string{fmt.Sprintf("LISTENINGPORT=%s", portopening2)}
		err = runContainerForProjectTestImage(client, image2, containernameproject_test, portopening2, inputEnv)
		if err != nil {
			log.Println(err)
		}

	} else if to_check_docker_image == 4 {

		fmt.Println("9) -------Create Image from a Docker File-------")

		// -- BUILD IMAGE FROM A DOCKER FILE
		client, err := client.NewEnvClient()
		if err != nil {
			log.Fatalf("Unable to create docker client: %s", err)
		}
		// Client, imagename and Dockerfile location

		tags := []string{"again_python"}

		dockerfile := "Dockerfile"

		err = buildImageforDockerFile1(client, tags, dockerfile)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

	} else if to_check_docker_image == 5 {

		fmt.Println("10) -------Create Image from a Docker File-------")

		// -- BUILD IMAGE FROM A DOCKER FILE
		client, err := client.NewEnvClient()
		if err != nil {
			log.Fatalf("Unable to create docker client: %s", err)
		}
		// Client, imagename and Dockerfile location

		tags := []string{"go_check"}

		dockerfile := "Dockerfile"

		err = buildImageforDockerFile2(client, tags, dockerfile)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

	} else {

		fmt.Println("6) -------Create Image from a Docker File-------")

		// -- BUILD IMAGE FROM A DOCKER FILE
		client, err := client.NewEnvClient()
		if err != nil {
			log.Fatalf("Unable to create docker client: %s", err)
		}
		// Client, imagename and Dockerfile location

		tags := []string{"python_hello_world_image"}
		dockerfile := `C:\Users\abdul\Desktop\T-Hex\t-hex\app\backend\PythonTest\Dockerfile`

		// dockerfile := `C:\Users\abdul\Desktop\go-workspace\src\ProjectTest\Dockerfile`
		err = buildImageforDockerFile(client, tags, dockerfile)
		if err != nil {
			log.Println(err)
		}

		// --
	}

}
