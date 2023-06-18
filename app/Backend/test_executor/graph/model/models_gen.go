package model

type Projectinfos struct {
	ID                  int    `json:"id"`
	GitEmail            string `json:"gitEmail"`
	GitProjectName      string `json:"gitProjectName"`
	GitLanguage         string `json:"gitLanguage"`
	GitNoOfTestCases    int    `json:"gitNoOfTestCases"`
	GitTestCaseFileName string `json:"gitTestCaseFileName"`
}

type Settings struct {
	ID                    int    `json:"id"`
	Browser               string `json:"browser"`
	Version               string `json:"version"`
	StepByStepDebugging   bool   `json:"stepByStepDebugging"`
	EnableLogs            bool   `json:"enableLogs"`
	Parallelism           bool   `json:"parallelism"`
	NumberOfParallelTests int    `json:"numberOfParallelTests"`
}


type CloneGitRepInput struct {
	GitRepPath string `json:"gitRepPath"`
}

type CreateDockerFileInput struct {
	DockerCmnds string `json:"dockerCmnds"`
}

type ProjectBuildImageInput struct {
	ImgName     string `json:"imgName"`
	ProjectPath string `json:"projectPath"`
}

type ProjectContainerEndInput struct {
	ContainerName string `json:"containerName"`
}

type ProjectContainerStartInput struct {
	ImageName     string `json:"imageName"`
	ContainerName string `json:"containerName"`
}

type SaveTestResultInput struct {
	Username string `json:"username"`
	Project  string `json:"project"`
	Type     string `json:"type"`
	Urlid    string `json:"urlid"`
	Testfile string `json:"testfile"`
	Status   string `json:"status"`
	Duration string `json:"duration"`
}

type SeleniumInput struct {
	ContainerName string `json:"containerName"`
}

type StartTestInput struct {
	SeleniumContainerName string `json:"seleniumContainerName"`
	ProjectPath           string `json:"projectPath"`
	ProjectImageName      string `json:"projectImageName"`
	ProjectContainerName  string `json:"projectContainerName"`
}

type UpdateTestInput struct {
	ID       int    `json:"id"`
	Status   string `json:"status"`
	TestPath string `json:"testPath"`
}
