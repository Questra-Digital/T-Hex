package graph

import (
	"context"
	"fmt"
	"log"
	"os"
	"test_executor/graph/model"
	f "test_executor/service"
	"time"

	"github.com/docker/docker/client"
	git "github.com/go-git/go-git/v5"
)

/*Delete All the Records*/

func (r *mutationResolver) DestroyAll(ctx context.Context) (string, error) {
	err := f.DeleteAllTests(ctx)
	if err != nil {
		return "Failed", nil
	}
	return "Destroyed All Test", nil
}

// InitiateTest is the resolver for the initiateTest field.
func (r *mutationResolver) InitiateTest(ctx context.Context) (string, error) {

	proj, err := f.GetAllProjectInfos(ctx)
	if err != nil {
		return "", err
	}

	lastProj := proj[len(proj)-1]
	fmt.Println("\nLast Record in Projectinfos Table: \n", lastProj)

	ID := lastProj.ID
	GitEmail := lastProj.GitEmail
	GitProjectName := lastProj.GitProjectName
	GitLanguage := lastProj.GitLanguage
	GitNoOfTestCases := lastProj.GitNoOfTestCases
	GitTestCaseFileName := lastProj.GitTestCaseFileName


	return f.InitiateTest(ctx, ID, GitEmail, GitProjectName, GitLanguage, GitNoOfTestCases, GitTestCaseFileName)
}

// StartTest is the resolver for the startTest field.
func (r *mutationResolver) StartTest(ctx context.Context, input model.StartTestInput) (string, error) {
	
	result, err := r.SeleniumPull(ctx)
	if err != nil {
		return "Failed", nil
	}
	result, err = r.SeleniumStart(ctx, input)
	if err != nil {
		return "Failed", nil
	}
	time.Sleep(8 * time.Second)
	result, err = r.ProjectBuildImage(ctx, input)
	if err != nil {
		return "Failed", nil
	}

	result, err = r.ProjectContainerStart(ctx, input)
	if err != nil {
		return "Failed", nil
	}
	return result, nil
}

// SaveTestResult is the resolver for the saveTestResult field.
func (r *mutationResolver) SaveTestResult(ctx context.Context, input model.SaveTestResultInput) (*model.Test, error) {
	fmt.Printf("saveTestResult Mutation\n")
	return f.SaveTestResult(ctx, input)
}

// UpdateTest is the resolver for the updateTest field.
func (r *mutationResolver) UpdateTest(ctx context.Context, input model.UpdateTestInput) (*model.Test, error) {
	return f.UpdateTestByID(ctx, input)
}

// SeleniumPull is the resolver for the seleniumPull field.
func (r *mutationResolver) SeleniumPull(ctx context.Context) (string, error) {
	fmt.Println("-------Pulling selenium image from docker hub registory-------")

	client, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("UNABLE to create docker client: %s", err)
	}

	image := "selenium/standalone-chrome"
	err = f.PullImageSelenium(client, image)

	if err != nil {
		log.Println(err)
		return "FAILED to Pull Selenium", nil
	}
	return "Selenium Pulled SUCCESSFULLY", nil
}

// SeleniumStart is the resolver for the seleniumStart field.
func (r *mutationResolver) SeleniumStart(ctx context.Context, input model.StartTestInput) (string, error) {
	fmt.Println("-------Running Selenium_Container-------")
	client, err := client.NewEnvClient()
	containername := input.SeleniumContainerName
	portopening := "4444"
	inputEnv1 := []string{fmt.Sprintf("LISTENINGPORT=%s", portopening)}
	image := "selenium/standalone-chrome"

	err = f.RunContainerSelenium(client, image, containername, portopening, inputEnv1)
	if err != nil {
		log.Println(err)
		return err.Error(), nil
	}
	return "Selenium Started SUCCESSFULLY", nil

}

// SeleniumStop is the resolver for the seleniumStop field.
func (r *mutationResolver) SeleniumStop(ctx context.Context, input model.SeleniumInput) (string, error) {
	fmt.Println("-------Stop Selenium_Container-------")
	client, err := client.NewEnvClient()
	err = f.StopsSeleniumContainer(client, input.ContainerName)
	if err != nil {
		log.Println(err)
		return "Selenium FAILED to Stop", nil
	}
	log.Println(err)
	return "Selenium Stopped SUCCESSFULLY", nil
}

// SeleniumStart is the resolver for the seleniumStart field.
func (r *mutationResolver) SeleniumRemove(ctx context.Context, input model.SeleniumInput) (string, error) {
	fmt.Println("-------(Removing Selenium Container)-------")

	client, err := client.NewEnvClient()
	err = f.RemoveSeleniumContainer(client, input.ContainerName)

	if err != nil {
		log.Println(err)
		return "UNABLE to remove Selenium ", nil
	}
	return "Selenium Removed SUCCESSFULLY", nil
}

// CreateDockerFile is the resolver for the createDockerFile field.
func (r *mutationResolver) CreateDockerFile(ctx context.Context, input model.CreateDockerFileInput) (string, error) {
	res, err := f.CreateDockerFile(ctx, input)
	if err != nil {
		log.Println(err)
		return "UNABLE to Create Docker File ", nil
	}
	fmt.Printf(res)
	return "Docker File Created SUCCESSFULLY", nil
}

// CloneGitRepository is the resolver for the cloneGitRepository field.
func (r *mutationResolver) CloneGitRepository(ctx context.Context, input model.CloneGitRepInput) (string, error) {

	/**/

	proj, err := f.GetAllProjectInfos(ctx)
	if err != nil {
		return "", err
	}

	lastProj := proj[len(proj)-1]
	fmt.Println("\nLast Record in Projectinfos Table: \n", lastProj)

	username := lastProj.GitEmail
	projectname := lastProj.GitProjectName

	path := input.GitRepPath

	basePath := "example/" + username + "/" + projectname + "/"
	fmt.Println(path)
	retStatement := "Clone Rep Successfull: " + path

	ctx = context.Background()
	url := path

	repoDir := basePath
	if err := os.MkdirAll(repoDir, os.ModePerm); err != nil {
		fmt.Println(err)
		return "MkdirALL Error", nil
	}

	repo, err := git.PlainCloneContext(ctx, repoDir, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

	if err != nil {
		fmt.Println(err)
		return "PlainCloneContext Error", nil
	}

	headRef, err := repo.Head()
	if err != nil {
		fmt.Println(err)
		return "Head() Error", nil
	}

	fmt.Println(headRef.Hash())

	return retStatement, nil
}

// ProjectBuildImage is the resolver for the projectBuildImage field.
func (r *mutationResolver) ProjectBuildImage(ctx context.Context, input model.StartTestInput) (string, error) {
	fmt.Println("-------Create Project Image from a Docker File-------")
	client, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("UNABLE to create docker client: %s", err)
		return "FAILED to create Docker Client", err
	}

	tags1 := []string{input.ProjectImageName}
	dockerfile1 := "Dockerfile"
	contextPath1 := input.ProjectPath
	err = f.BuildImageforDockerFile2(client, tags1, dockerfile1, contextPath1)
	if err != nil {
		log.Println(err)
		os.Exit(1)
		return "FAILED to Build Project Image", nil
	}
	return "Project Image Built SUCCESSFULLY", nil
}

// ProjectContainerStart is the resolver for the projectContainerStart field.
func (r *mutationResolver) ProjectContainerStart(ctx context.Context, input model.StartTestInput) (string, error) {
	fmt.Println("5) -------Running Project Contianer-------")
	client, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("UNABLE to create docker client: %s", err)
		return "FAILED to create Docker Client", err
	}

	containernameproject_test1 := input.ProjectContainerName 
	portopening2 := "8085"
	image2 := input.ProjectImageName 
	inputEnv2 := []string{fmt.Sprintf("LISTENINGPORT=%s", portopening2)}
	err = f.RunContainerForProjectTestImage(client, image2, containernameproject_test1, portopening2, inputEnv2)
	if err != nil {
		log.Println(err)
	}
	return "Project Container Started Successfully", nil
}

// ProjectContainerStop is the resolver for the projectContainerStop field.
func (r *mutationResolver) ProjectContainerStop(ctx context.Context, input model.ProjectContainerEndInput) (string, error) {
	fmt.Println("-------Stopping Project Container-------")
	client, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("UNABLE to create docker client: %s", err)
		return "FAILED to create Docker Client", err
	}

	err = f.StopProjectContainer(client, input.ContainerName)
	if err != nil {
		log.Println(err)
		return "Project Container FAILED to Stop", nil
	}
	return "Project Container Stopped SUCCESSFULLY", nil
}

// ProjectContainerRemove is the resolver for the projectContainerRemove field.
func (r *mutationResolver) ProjectContainerRemove(ctx context.Context, input model.ProjectContainerEndInput) (string, error) {
	fmt.Println("-------Removing Project Container-------")
	client, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("UNABLE to create docker client: %s", err)
		return "FAILED to create Docker Client", err
	}

	err = f.RemoveProjectContainer(client, input.ContainerName)
	if err != nil {
		log.Println(err)
		return err.Error(), nil
	}
	return "Project Container Removed SUCCESSFULLY", nil
}

// Tests is the resolver for the tests field.
func (r *queryResolver) Tests(ctx context.Context) ([]*model.Test, error) {
	return f.GetAllTest(ctx)
}

// Test is the resolver for the test field.
func (r *queryResolver) Test(ctx context.Context, id int) (*model.Test, error) {
	return f.GetTestByID(ctx, id)
}

// Settings is the resolver for the settings field.
func (r *queryResolver) Settings(ctx context.Context) ([]*model.Settings, error) {
	return f.GetAllSettings(ctx)
}

// Setting is the resolver for the setting field.
func (r *queryResolver) Setting(ctx context.Context, id int) (*model.Settings, error) {
	panic(fmt.Errorf("not implemented: Setting - setting"))
}

// ProjectInfos is the resolver for the projectInfos field.
func (r *queryResolver) ProjectInfos(ctx context.Context) ([]*model.Projectinfos, error) {
	return f.GetAllProjectInfos(ctx)
}

// ProjectInfo is the resolver for the projectInfo field.
func (r *queryResolver) ProjectInfo(ctx context.Context, id int) (*model.Projectinfos, error) {
	panic(fmt.Errorf("not implemented: ProjectInfo - projectInfo"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
