
package model

type GetSettingInput struct {
	Browser             string `json:"browser"`
	Version             string `json:"version"`
	StepByStepDebugging bool   `json:"stepByStepDebugging"`
	EnableLogs          bool   `json:"enableLogs"`
	Parallelism         bool   `json:"parallelism"`
}

type ProjectInfoInput struct {
	GitEmail            string `json:"gitEmail"`
	GitProjectName      string `json:"gitProjectName"`
	GitLanguage         string `json:"gitLanguage"`
	GitNoOfTestCases    int    `json:"gitNoOfTestCases"`
	GitTestCaseFileName string `json:"gitTestCaseFileName"`
}
