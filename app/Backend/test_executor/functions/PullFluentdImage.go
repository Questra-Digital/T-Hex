package functions

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func PullFluentdImage(client *client.Client) error {
	ctx := context.Background()
	image := "fluent/fluentd:v1.13"
	reader, err := client.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, reader)
	return nil
}
