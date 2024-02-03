package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"test_executor/configdb"
	"test_executor/graph/model"
	"time"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/drive/v3"

	f "test_executor/functions"

	"github.com/docker/docker/client"

	"gorm.io/gorm"
)

func DeleteAllTests(ctx context.Context) error {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	if err := db.Exec("DELETE FROM tests").Error; err != nil {
		return err
	}

	client, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("Unable to create docker client: %s", err)
	}

	proj, err := GetAllProjectInfos(ctx)
	if err != nil {
		return err
	}
	lastProj := proj[len(proj)-1]
	fmt.Println("\nLast Record in Projectinfos Table: \n", lastProj)
	fmt.Println(lastProj.ID)

	username := lastProj.GitEmail
	filename := lastProj.GitTestCaseFileName
	nooftestcases := lastProj.GitNoOfTestCases
	language := lastProj.GitLanguage

	basePath := "example/" + username + "/"
	folderPath := basePath

	err = os.RemoveAll(folderPath)
	if err != nil {
		fmt.Println("Failed to remove the folder:", err)
		return err
	}

	fmt.Println("Folder removed successfully.")

	fmt.Println("Stop and Removing the fluentd containe")

	for i := 1; i < nooftestcases+1; i++ {

		containername_f := "Project_" + strconv.Itoa(i) + "_Fluentd"
		err = f.StopAndRemoveFluentdContainer(client, containername_f)

		if err != nil {

			log.Println(err)
			return err
		}
	}

	fmt.Println("Stop and Removing the Project containe")

	for i := 1; i < nooftestcases+1; i++ {

		ContainerName := strings.ToLower(language) + "-" + "test" + "-" + "container" + "-" + strconv.Itoa(i)
		err = f.StopAndRemoveProjectContainer(client, ContainerName)

		if err != nil {

			log.Println(err)
			return err
		}
	}

	fmt.Println("Stop and Removing the Video containe")

	for i := 1; i < nooftestcases+1; i++ {

		Container_Video_name := filename + "-test-video-" + strconv.Itoa(i)
		err = f.StopsAndRemoveVideoContainer(client, Container_Video_name)

		if err != nil {

			log.Println(err)
			return err
		}
	}

	fmt.Println("Stop and Removing the Selenium containe")

	Container_Selenium_name := "Selenium_Container"

	err = f.StopsAndRemoveSeleniumContainer(client, Container_Selenium_name)

	if err != nil {

		log.Println(err)
		return err
	}

	return nil
}

func SaveTestResult(ctx context.Context, input model.SaveTestResultInput) (*model.Test, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	tests := model.Test{
		Username: input.Username,
		Project:  input.Project,
		Type:     input.Type,
		Urlid:    input.Urlid,
		Testfile: input.Testfile,
		Status:   input.Status,
		Duration: input.Duration,
	}

	if err := db.Table("tests").Create(&tests).Error; err != nil {
		return nil, err
	}
	return &tests, nil
}

func InitiateTest(ctx context.Context, ID int, GitEmail string, GitProjectName string, GitLanguage string, GitNoOfTestCases int, GitTestCaseFileName string) (string, error) {

	fmt.Println(ID)
	fmt.Println(GitEmail)
	fmt.Println(GitProjectName)
	fmt.Println(GitLanguage)
	fmt.Println(GitNoOfTestCases)
	fmt.Println(GitTestCaseFileName)

	client, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("Unable to create docker client: %s", err)
	}

	username := GitEmail
	projectname := GitProjectName
	filename := GitTestCaseFileName
	nooftestcases := GitNoOfTestCases
	language := GitLanguage

	baseDir := "C:/Users/abdul/Desktop/Work By Waleed FYP/sideby_fyp_project/"
	basePath := "example/" + username + "/" + projectname + "/"

	start_time := make([]time.Time, nooftestcases)
	end_time := make([]string, nooftestcases)

	test_result := make([]string, nooftestcases)

	// Fluentd config

	fluentconfigcontent := `

	<source>
	  @type forward
	  port 24224
	</source>
		`

	for i := 1; i < nooftestcases+1; i++ {

		content2 := fluentconfigcontent
		content2 += "\n"
		content2 += "<match "
		content2 += "fluent_d_container_"
		content2 += strconv.Itoa(i)
		content2 += ">"
		content2 += "\n"
		content2 += "  @type file"
		content2 += "\n"
		content2 += "  path /fluentd/log/fluentdlogs"
		content2 += "\n</match>"

		basePath := "example/" + username + "/" + projectname + "/"
		configfile := basePath + "container-logs-"
		configfile += strconv.Itoa(i)
		configfile += ".conf"

		err := ioutil.WriteFile(configfile, []byte(content2), 0644)
		if err != nil {
			fmt.Println("Error creating configfile:", err)
			ioutil.WriteFile(configfile, []byte(content2), 0644)
		}
	}

	// Pull Images

	err = f.PullImageSelenium(client)
	if err != nil {

		log.Println(err) ///
		return "Pull Selenium Failed", err
	}

	err = f.PullFluentdImage(client)
	if err != nil {

		log.Println(err) ///
		return "Pull Fluentd Failed", err
	}

	err = f.PullImageVideo(client)
	if err != nil {

		log.Println(err) ///
		return "Pull Video Failed", err
	}

	// Run Selenium Container

	fmt.Println("2) -------Running Selenium_Container-------")
	image_s := "selenium/standalone-chrome"
	containername_s := "Selenium_Container"
	portopening := "4444"
	inputEnv1 := []string{fmt.Sprintf("LISTENINGPORT=%s", portopening)}
	err = f.RunContainerSelenium(client, image_s, containername_s, portopening, inputEnv1)
	if err != nil {
		log.Println(err)
	}

	// Build Docker Image from dockerfile-contents

	for i := 1; i < nooftestcases+1; i++ {

		fmt.Println("p2-task) -------Create " + language + " Test Image from a Docker File" + strconv.Itoa(i) + "-------")

		tag_string := strings.ToLower(language) + "-" + "test" + "-" + "image" + "-" + strconv.Itoa(i)
		tags1 := []string{tag_string}
		dockerfile1 := "Dockerfile"
		dockerfile1 += strconv.Itoa(i)
		contextPath1 := "example\\" + username + "\\" + projectname
		err = f.BuildImageforDockerFile2(client, tags1, dockerfile1, contextPath1)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

	}

	// Run fluentd Container.

	for i := 1; i < nooftestcases+1; i++ {
		image_f := "fluent/fluentd:v1.13"
		containername_f := "Project_" + strconv.Itoa(i) + "_Fluentd"
		err = f.RunProjectFluentd(client, image_f, containername_f, strconv.Itoa(i), baseDir, basePath)
		if err != nil {

			log.Println(err) ///
			return "FLuentf Running Failed", err
		}
	}

	// Runs Video Container

	for i := 1; i < nooftestcases+1; i++ {

		video_name := "video-" + filename + "-" + strconv.Itoa(i)
		Container_Video_name := filename + "-test-video-" + strconv.Itoa(i)

		imagev := "selenium/video"
		err = f.RunVideoContainer(client, imagev, Container_Video_name, video_name, baseDir, basePath)

		if err != nil {
			log.Println(err)
		}
	}

	// Run Project Container.

	for i := 1; i < nooftestcases+1; i++ {

		fmt.Println("p3-task) -------Run " + language + " Test Container from " + strings.ToLower(language) + "-test-image-" + strconv.Itoa(i) + "-------")

		ImageName := strings.ToLower(language) + "-" + "test" + "-" + "image" + "-" + strconv.Itoa(i)
		ContainerName := strings.ToLower(language) + "-" + "test" + "-" + "container" + "-" + strconv.Itoa(i)

		Containerport := "8080"
		string_c, _ := strconv.Atoi(Containerport)
		Containerport = strconv.Itoa(string_c + i)

		inputEnvImage := []string{fmt.Sprintf("LISTENINGPORT=%s", Containerport)}

		containerID := "Project_" + strconv.Itoa(i) + "_Fluentd"

		inspect, err := client.ContainerInspect(context.Background(), containerID)
		if err != nil {
			panic(err)
		}

		ipAddress := inspect.NetworkSettings.Networks["bridge"].IPAddress
		fmt.Println("Container IP address:", ipAddress)

		start_time[i-1] = time.Now()
		err = f.RunContainerForProjectTestImage(client, ImageName, ContainerName, Containerport, inputEnvImage, ipAddress, strconv.Itoa(i), baseDir, basePath)

		if err != nil {
			log.Println(err)
		}

	}

	fmt.Println("**** Before **** making the call for the container to Stop")

	// Stops Video Container

	for i := 1; i < nooftestcases+1; i++ {

		containerID := strings.ToLower(language) + "-test-container-" + strconv.Itoa(i)

		end_time[i-1] = time.Since(start_time[i-1]).String()

		cmd := exec.Command("docker", "wait", containerID)

		output, err := cmd.Output()
		if err != nil {
			panic(err)
		}

		exitCode := strings.TrimSpace(string(output))

		fmt.Printf("Container exited with status code: %s\n", exitCode)

		if exitCode == "1" {
			fmt.Printf("containerID %s Failed", containerID)
			test_result[i-1] = "failed"
		} else if exitCode == "0" {
			fmt.Printf("containerID %s passed", containerID)
			test_result[i-1] = "passed"
		}

		fmt.Println("p3-task) -------Stop " + language + " Video Container")

		Container_Video_name := filename + "-test-video-" + strconv.Itoa(i)

		fmt.Println(" -------Stop Video_Container-------")
		f.StopsVideoContainer(client, Container_Video_name)
	}

	fmt.Println("---- After ---- making the call for the container to Stop")

	for i := 0; i < nooftestcases; i++ {
		fmt.Printf("The time for %d, is %s:\n", i, end_time[i])
	}

	functionToPushToDrive(username, projectname, nooftestcases, filename, end_time, test_result)

	return "Ran All The Fucnitons", nil
}

func ServiceAccount(secretFile string) *http.Client {
	b, err := ioutil.ReadFile(secretFile)
	if err != nil {
		log.Fatal("error while reading the credential file", err)
	}
	var s = struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}{}
	json.Unmarshal(b, &s)
	config := &jwt.Config{
		Email:      s.Email,
		PrivateKey: []byte(s.PrivateKey),
		Scopes: []string{
			drive.DriveScope,
		},
		TokenURL: google.JWTTokenURL,
	}
	client := config.Client(context.Background())
	return client
}

func createFile(service *drive.Service, name string, mimeType string, content io.Reader, parentId string) (*drive.File, error) {
	f := &drive.File{
		MimeType: mimeType,
		Name:     name,
		Parents:  []string{parentId},
	}
	file, err := service.Files.Create(f).Media(content).Do()

	if err != nil {
		log.Println("Could not create file: " + err.Error())
		return nil, err
	}

	return file, nil
}

type appendData struct {
	testID int
	name   string
}

func functionToPushToDrive(username string, projectname string, nooftestcases int, testfile string, duration []string, test_result []string) (string, error) {

	db := configdb.ConnectCockroachDB()

	for i := 0; i < nooftestcases; i++ {

		basePath := "example/" + username + "/" + projectname + "/" + "fluentd" + strconv.Itoa(i+1)

		filePrefix := "buffer"

		files, err := ioutil.ReadDir(basePath)
		if err != nil {
			fmt.Println("Error reading directory:", err)
			return "Pull Successfull", nil
		}

		for _, file := range files {
			fileName := file.Name()
			if file.IsDir() {
				continue
			}

			if filepath.HasPrefix(fileName, filePrefix) {
				if strings.HasSuffix(fileName, ".log") {

					testfolder := basePath + "/"

					f, err := os.Open(testfolder + fileName)

					if err != nil {
						panic(fmt.Sprintf("cannot open file: %v", err))
					}

					defer f.Close()

					client := ServiceAccount("Client Secret/client_secret.json")

					srv, err := drive.New(client)
					if err != nil {
						log.Fatalf("Unable to retrieve drive Client %v", err)
					}

					folderId := "1Z4_EPT_ZF62pRHrkmtyJdAkoBL0eeHo9"

					SavePath := fileName
					file, err := createFile(srv, SavePath, "application/octet-stream", f, folderId)

					if err != nil {
						panic(fmt.Sprintf("Could not create file: %v\n", err))
					}

					fmt.Printf("File '%s' successfully uploaded", file.Name)
					fmt.Printf("\nFile Id: '%s' ", file.Id)

					fmt.Printf("\nusername %s, project %s, type %s, URLid %s, testfile %s, duration %s\n",
						username, projectname, "log", file.Id, testfile+strconv.Itoa(i+1), duration[i])

					_, err = SaveResult(db, username, projectname, "log", file.Id, testfile+strconv.Itoa(i+1), duration[i], test_result[i])

					if err != nil {
						return "Error", err
					}
				}
			}
		}

	}

	for i := 0; i < nooftestcases; i++ {

		basePath := "example/" + username + "/" + projectname + "/" + "Screenshots" + strconv.Itoa(i+1)

		files, err := ioutil.ReadDir(basePath)
		if err != nil {
			fmt.Println("Error reading directory:", err)
		}

		for _, file := range files {
			fileName := file.Name()
			if file.IsDir() {
				continue
			}

			if strings.HasSuffix(fileName, ".png") {

				testfolder := basePath + "/"

				f, err := os.Open(testfolder + fileName)

				if err != nil {
					panic(fmt.Sprintf("cannot open file: %v", err))
				}

				defer f.Close()

				client := ServiceAccount("Client Secret/client_secret.json")

				srv, err := drive.New(client)
				if err != nil {
					log.Fatalf("Unable to retrieve drive Client %v", err)
				}

				folderId := "1Z4_EPT_ZF62pRHrkmtyJdAkoBL0eeHo9"

				SavePath := fileName
				file, err := createFile(srv, SavePath, "application/octet-stream", f, folderId)

				if err != nil {
					panic(fmt.Sprintf("Could not create file: %v\n", err))
				}

				fmt.Printf("File '%s' successfully uploaded", file.Name)
				fmt.Printf("\nFile Id: '%s' ", file.Id)

				fmt.Printf("\nusername %s, project %s, type %s, URLid %s, testfile %s, duration %s\n",
					username, projectname, "screenshot", file.Id, testfile+strconv.Itoa(i+1), duration[i])

				_, err = SaveResult(db, username, projectname, "screenshot", file.Id, testfile+strconv.Itoa(i+1), duration[i], test_result[i])

				if err != nil {
					return "Error", err
				}

			}

		}

	}

	/* For Video's*/
	for i := 0; i < nooftestcases; i++ {

		testfolder := "example/" + username + "/" + projectname + "/" + "video-" + testfile + "-" + strconv.Itoa(i+1) + ".mp4/"

		f, err := os.Open(testfolder + "video.mp4")

		if err != nil {
			panic(fmt.Sprintf("cannot open file: %v", err))
		}

		defer f.Close()

		client := ServiceAccount("Client Secret/client_secret.json")

		srv, err := drive.New(client)
		if err != nil {
			log.Fatalf("Unable to retrieve drive Client %v", err)
		}

		folderId := "1Z4_EPT_ZF62pRHrkmtyJdAkoBL0eeHo9"

		SavePath := username + "-" + testfile
		file, err := createFile(srv, SavePath, "application/octet-stream", f, folderId)

		if err != nil {
			panic(fmt.Sprintf("Could not create file: %v\n", err))
		}

		fmt.Printf("File '%s' successfully uploaded", file.Name)
		fmt.Printf("\nFile Id: '%s' ", file.Id)

		fmt.Printf("\nusername %s, project %s, type %s, URLid %s, testfile %s, duration %s\n",
			username, projectname, "video", file.Id, testfile+strconv.Itoa(i+1), duration[i])

		_, err = SaveResult(db, username, projectname, "video", file.Id, testfile+strconv.Itoa(i+1), duration[i], test_result[i])

		if err != nil {
			return "Error", err
		}

	}

	fmt.Println("Save Resuls ok!")

	return "Pull Successfull", nil

}
func SaveResult(db *gorm.DB, username string, projectname string, file_type string, url_id string, testfile string, duration string, testresult string) (*model.Test, error) {

	tests := model.Test{
		Username: username,
		Project:  projectname,
		Type:     file_type,
		Urlid:    url_id,
		Testfile: testfile,
		Status:   testresult,
		Duration: duration,
	}

	if err := db.Table("tests").Create(&tests).Error; err != nil {
		return nil, err
	}
	return &tests, nil
}

func CreateDockerFile(ctx context.Context, input model.CreateDockerFileInput) (string, error) {

	result := input.DockerCmnds

	proj, err := GetAllProjectInfos(ctx)
	if err != nil {
		return "GetAllProjectInfos Failed", err
	}
	lastProj := proj[len(proj)-1]
	fmt.Println("\nLast Record in Projectinfos Table: \n", lastProj)
	fmt.Println(lastProj.ID)

	username := lastProj.GitEmail
	projectname := lastProj.GitProjectName
	filename := lastProj.GitTestCaseFileName
	nooftestcases := lastProj.GitNoOfTestCases
	language := lastProj.GitLanguage

	dockerfilecontent := result
	if language == "Python" || language == "python" {

		for i := 1; i < nooftestcases+1; i++ {

			content1 := dockerfilecontent
			content1 += "\n"
			content1 += "COPY "
			content1 += filename
			content1 += strconv.Itoa(i)
			content1 += ".py /usr/app/src"
			content1 += "\n"

			content1 += `CMD ["python3", "./`
			content1 += filename
			content1 += strconv.Itoa(i)
			content1 += `.py"]`

			fmt.Println(content1)

			basePath := "example/" + username + "/" + projectname + "/"
			dockerfile := basePath + "Dockerfile"
			dockerfile += strconv.Itoa(i)
			fmt.Println(dockerfile)
			err := ioutil.WriteFile(dockerfile, []byte(content1), 0644)
			if err != nil {
				fmt.Println("Error creating dockerfile:", err)
				return "error", nil
			}
		}

	}

	return result, nil
}

func GetAllTest(ctx context.Context) ([]*model.Test, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var tests []*model.Test
	if err := db.Table("tests").Find(&tests).Error; err != nil {
		return nil, err
	}

	return tests, nil
}

func GetTestByID(ctx context.Context, id int) (*model.Test, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var tests model.Test
	if err := db.Table("tests").Where("id = ?", id).Take(&tests).Error; err != nil {
		return nil, err
	}

	return &tests, nil
}

func UpdateTestByID(ctx context.Context, input model.UpdateTestInput) (*model.Test, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Table("tests").Where("id = ?", input.ID).Update("status", input.Status).Error; err != nil {
		return nil, err
	}
	if err := db.Table("tests").Where("id = ?", input.ID).Update("test_path", input.TestPath).Error; err != nil {
		return nil, err
	}
	return GetTestByID(ctx, input.ID)
}

func GetAllSettings(ctx context.Context) ([]*model.Settings, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var settings []*model.Settings
	if err := db.Table("settings").Find(&settings).Error; err != nil {
		return nil, err
	}

	return settings, nil
}
func GetAllProjectInfos(ctx context.Context) ([]*model.Projectinfos, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var projectinfos []*model.Projectinfos
	if err := db.Table("projectinfos").Find(&projectinfos).Error; err != nil {
		return nil, err
	}

	return projectinfos, nil
}

func GetProjectInfoByID(ctx context.Context, id int) (*model.Projectinfos, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var projectinfos model.Projectinfos
	if err := db.Table("projectinfos").Where("id = ?", id).Take(&projectinfos).Error; err != nil {
		return nil, err
	}

	return &projectinfos, nil
}
func GetSettingByID(ctx context.Context, id int) (*model.Settings, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var settings model.Settings
	if err := db.Table("settings").Where("id = ?", id).Take(&settings).Error; err != nil {
		return nil, err
	}

	return &settings, nil
}
