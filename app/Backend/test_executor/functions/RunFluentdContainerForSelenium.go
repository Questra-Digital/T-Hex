package functions

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func RunFluentdContainer(client *client.Client, imagename string, containername string, baseDir string) error {

	hostConfig := &container.HostConfig{
		Binds: []string{baseDir + "app/Backend/test_executor/fluentd/frosty_hofstadter.conf:/fluentd/etc/fluent.conf",
			baseDir + "app/Backend/test_executor/fluentd:/fluentd/log/fluentdlogs"},
	}

	config := &container.Config{
		Image: imagename,
		Cmd:   []string{"fluentd", "-c", "/fluentd/etc/fluent.conf"},
	}

	cont, err := client.ContainerCreate(
		context.Background(),
		config,
		hostConfig,
		nil, 
		nil, 
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
