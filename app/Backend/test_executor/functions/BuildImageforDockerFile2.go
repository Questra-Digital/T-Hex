package functions

import (
	"archive/tar"
	"bytes"
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// BuildImageforDockerFile2 builds a Docker image from a Dockerfile and context path using Docker client
func BuildImageforDockerFile2(client *client.Client, tags []string, dockerfile string, contextPath string) error {
	ctx := context.Background()

	// Create a tar buffer from the context path
	tarBuffer, err := createDockerBuildContext(contextPath)
	if err != nil {
		return err
	}
	defer tarBuffer.Close()

	// Build options for the image
	buildOptions := types.ImageBuildOptions{
		Context:    tarBuffer,
		Dockerfile: dockerfile,
		Remove:     true,
		Tags:       tags,
	}

	// Initiate the image build process using the Docker client
	imageBuildResponse, err := client.ImageBuild(ctx, tarBuffer, buildOptions)
	if err != nil {
		return err
	}
	defer imageBuildResponse.Body.Close()

	// Copy the build output to the standard output (os.Stdout)
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		return err
	}

	return nil
}

// createDockerBuildContext creates a tar buffer for the Docker build context
func createDockerBuildContext(contextPath string) (*bytes.Reader, error) {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)

	// Set the current working directory to the context path
	cwd := contextPath

	// Walk through the context path and add files/directories to the tar writer
	err := filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create a tar header from file info
		header, err := tar.FileInfoHeader(info, path)
		if err != nil {
			return err
		}

		// Set the file name within the tar file
		header.Name = strings.TrimPrefix(path, cwd+string(filepath.Separator))

		// Write the header to the tar file
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// If it's not a directory, open the file and copy its contents to the tar writer
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
		return nil, err
	}

	// Return a reader for the Docker context tar
	return bytes.NewReader(buf.Bytes()), nil
}
