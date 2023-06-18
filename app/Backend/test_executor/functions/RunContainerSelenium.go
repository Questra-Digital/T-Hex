package functions

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	natting "github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

func RunContainerSelenium(client *client.Client, imagename string, containername string, port string, inputEnv []string) error {

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
