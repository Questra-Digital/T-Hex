// package functions

// import (
// 	"context"
// 	"log"

// 	"github.com/docker/docker/api/types"
// 	"github.com/docker/docker/api/types/container"
// 	"github.com/docker/docker/client"
// )

// func RunFluentdContainer(client *client.Client, imagename string, containername string) error {

// 	hostConfig := &container.HostConfig{
// 		// Binds: []string{"C:\\Users\\abdul\\Desktop\\T-Hex\\t-hex\\app\\Backend\\test_executor\\fluentd:/fluentd/etc"},
// 		// PortBindings: natting.PortMap{
// 		// 	"24224/tcp": []natting.PortBinding{
// 		// 		{
// 		// 			HostIP:   "",
// 		// 			HostPort: "24224",
// 		// 		},
// 		// 	},
// 		// },
// 		// AutoRemove: true,
// 	}

// 	config := &container.Config{
// 		Image: imagename,
// 		Cmd:   []string{"fluentd", "-c", "/fluentd/etc/fluent.conf"},
// 		// Tty:   true,
// 	}

// 	cont, err := client.ContainerCreate(
// 		context.Background(),
// 		config,
// 		hostConfig,
// 		nil, // network configuration not needed
// 		nil, // platform not needed
// 		containername,
// 	)

// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	if err := client.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{}); err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	log.Printf("Container %s is created", cont.ID)

// 	return nil
// }

package functions

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func RunProjectFluentd(client *client.Client, imagename string, containername string, i string, baseDir string, basepath string) error {

	hostConfig := &container.HostConfig{

		// C:/Users/abdul/Desktop/t-hex2/t-hex/app/Backend/test_executor/example/abdullahsaleem40404@gmail.com/FYP-Sample-Python-Test/container-logs-1.conf

		Binds: []string{baseDir + "app/Backend/test_executor/" + basepath + "container-logs-" + i + ".conf:/fluentd/etc/fluent.conf",
			baseDir + "app/Backend/test_executor/" + basepath + "fluentd" + i + ":/fluentd/log/fluentdlogs"},

		// PortBindings: natting.PortMap{
		// 	"24229/tcp": []natting.PortBinding{
		// 		{
		// 			HostIP:   "0.0.0.0",
		// 			HostPort: "24229",
		// 		},
		// 	},
		// },
		// AutoRemove: false,
	}

	config := &container.Config{
		Image: imagename,
		Cmd:   []string{"fluentd", "-c", "/fluentd/etc/fluent.conf"},
		// Tty:   true,
	}

	cont, err := client.ContainerCreate(
		context.Background(),
		config,
		hostConfig,
		nil, // network configuration not needed
		nil, // platform not needed
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
