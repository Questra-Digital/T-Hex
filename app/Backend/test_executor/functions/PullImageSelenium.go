package functions

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

<<<<<<< HEAD
func PullImageSeleniumForChrome(client *client.Client) error {
=======
func PullImageSelenium(client *client.Client) error {
>>>>>>> 85a3ff8396edf15420149ecce8110aaa01c0dac9
	ctx := context.Background()
	image := "selenium/standalone-chrome"
	reader, err := client.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, reader)
	return nil
}
<<<<<<< HEAD

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
=======
>>>>>>> 85a3ff8396edf15420149ecce8110aaa01c0dac9
