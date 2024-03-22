package functions

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	natting "github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

// ContainerInfo represents information about a Docker container
type ContainerInfo struct {
	ImageName      string
	ContainerName  string
	Port           string
	InputEnv       []string
	ContainerIP    string
	ContainerIndex string
	BaseDir        string
	BasePath       string
}

func RunContainerForProjectTestImageForParralel(client *client.Client, imagename string, containername string, port string, inputEnv []string, containerip string, i string, baseDir string, basepath string, wg *sync.WaitGroup) {
	defer wg.Done()

	newport, err := natting.NewPort("tcp", port)
	if err != nil {
		log.Println("Unable to create docker port:", err)
		return
	}
	platform := &v1.Platform{}

	hostConfig := &container.HostConfig{
		Links: []string{"Selenium_Container:alias"},
		Binds: []string{baseDir + "app/Backend/test_executor/" + basepath + "Screenshots" + i + ":/usr/app/src/T-HEX"},

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
			Type: "fluentd",
			Config: map[string]string{
				"fluentd-address": containerip + ":24224",
				"tag":             "fluent_d_container_" + i,
			},
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
		return
	}

	client.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})
	log.Printf("Container %s is created", cont.ID)
}

// RunContainerConcurrently runs existing Docker containers concurrently
func RunContainerConcurrently(client *client.Client, containers []ContainerInfo) {
	var wg sync.WaitGroup

	// Run containers concurrently using goroutines
	for _, container := range containers {
		wg.Add(1)
		go RunContainerForProjectTestImageForParralel(client, container.ImageName, container.ContainerName, container.Port, container.InputEnv, container.ContainerIP, container.ContainerIndex, container.BaseDir, container.BasePath, &wg)
	}

	// Wait for all containers to finish
	wg.Wait()
}
