package functions

import (
	"context"
	"log"

	"github.com/docker/docker/client"
)

func StopsVideoContainer(client *client.Client, containername string) error {
	ctx := context.Background()

	if err := client.ContainerStop(ctx, containername, nil); err != nil {
		log.Printf("Unable to stop container %s: %s", containername, err)
	}

	return nil
}
