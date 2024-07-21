package functions

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

func RunVideoContainer(client *client.Client, imagename string, containername string, video_name string, baseDir string, basepath string) error {

	hostConfig := &container.HostConfig{
		Links: []string{"Selenium_Container:selenium"},
		Binds: []string{baseDir + "app/Backend/test_executor/" + basepath + "/" + video_name + ".mp4" + ":/videos"},
	}

	config := &container.Config{
		Image:    imagename,
		Hostname: fmt.Sprintf("%s-hostnameexample", imagename),

	}

	platform := &v1.Platform{}

	cont, err := client.ContainerCreate(
		context.Background(),
		config,
		hostConfig,
		nil, 
		platform,
		containername,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	if err := client.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{}); err != nil {
		log.Println(err)
		return err
	}

	log.Printf("Container %s is created", cont.ID)

	return nil
}
