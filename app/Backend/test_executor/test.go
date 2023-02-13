package main

import (
	"fmt"
	"log"
	"os"
	"time"

	f "test_executor/functions"

	"github.com/docker/docker/client"
)

func main() {

	fmt.Println("1) -------Pulling selenium image from docker hub registory-------")

	client, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("Unable to create docker client: %s", err)
	}

	image := "selenium/standalone-chrome"
	err = f.PullImageSelenium(client, image)
	if err != nil {

		log.Println(err)
		return
	}

	fmt.Println("2) -------Running Selenium_Container-------")

	containername := "Selenium_Container"
	portopening := "4444"
	inputEnv1 := []string{fmt.Sprintf("LISTENINGPORT=%s", portopening)}
	err = f.RunContainerSelenium(client, image, containername, portopening, inputEnv1)
	if err != nil {
		log.Println(err)
	}

	time.Sleep(10 * time.Second)

	fmt.Println("3) -------Create PythonTest Image from a Docker File-------")

	tags1 := []string{"python_test"}
	dockerfile1 := "Dockerfile"
	// contextPath1 := "C:\\Users\\abdul\\Desktop\\T-Hex\\t-hex\\app\\backend\\PythonTest"
	contextPath1 := "example\\PythonTest"
	err = f.BuildImageforDockerFile2(client, tags1, dockerfile1, contextPath1)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println("4) -------Create GoTest Image from a Docker File-------")

	tags2 := []string{"go_test"}
	dockerfile2 := "Dockerfile"
	//contextPath2 := "C:\\Users\\abdul\\Desktop\\T-Hex\\t-hex\\app\\backend\\GoTest"
	contextPath2 := "example\\GoTest"
	err = f.BuildImageforDockerFile2(client, tags2, dockerfile2, contextPath2)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	time.Sleep(10 * time.Second)

	fmt.Println("5) -------Running PythonTest-------")

	containernameproject_test1 := "PythonTest_Container"
	portopening2 := "8080"
	image2 := "python_test"
	inputEnv2 := []string{fmt.Sprintf("LISTENINGPORT=%s", portopening2)}
	err = f.RunContainerForProjectTestImage(client, image2, containernameproject_test1, portopening2, inputEnv2)
	if err != nil {
		log.Println(err)
	}

	time.Sleep(10 * time.Second)

	fmt.Println("6) -------Running GoTest-------")

	containernameproject_test2 := "GoTest_Container"
	portopening3 := "8080"
	image3 := "go_test"
	inputEnv3 := []string{fmt.Sprintf("LISTENINGPORT=%s", portopening3)}
	err = f.RunContainerForProjectTestImage(client, image3, containernameproject_test2, portopening3, inputEnv3)
	if err != nil {
		log.Println(err)
	}

	time.Sleep(40 * time.Second)

	fmt.Println("7) ------- ~(Stop and removing)~ PythonTest_Container-------")

	f.StopAndRemoveProjectContainer(client, "PythonTest_Container")

	time.Sleep(20 * time.Second)

	fmt.Println("8) ------- ~(Stop and removing)~ GoTest_Container-------")

	f.StopAndRemoveProjectContainer(client, "GoTest_Container")

	time.Sleep(10 * time.Second)

	fmt.Println("9) -------Stop Selenium_Container-------")
	f.StopsSeleniumContainer(client, "Selenium_Container")

}
