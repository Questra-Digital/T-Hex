package model

type Projectinfos struct {
	ID                  int    `json:"id"`
	GitEmail            string `json:"gitEmail"`
	GitProjectName      string `json:"gitProjectName"`
	GitLanguage         string `json:"gitLanguage"`
	GitNoOfTestCases    int    `json:"gitNoOfTestCases"`
	GitTestCaseFileName string `json:"gitTestCaseFileName"`
}
