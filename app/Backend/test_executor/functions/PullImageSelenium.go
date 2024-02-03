package functions

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func PullImageSeleniumForChrome(client *client.Client) error {
	ctx := context.Background()
	image := "selenium/standalone-chrome"
	reader, err := client.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, reader)
	return nil
}

func PullImageSeleniumForEdge(client *client.Client) error {
	ctx := context.Background()
	image := "selenium/standalone-edge" // Update the image to use Selenium with Edge
	reader, err := client.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, reader)
	return nil
}
