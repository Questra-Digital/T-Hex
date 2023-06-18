package model

type Test struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Project  string `json:"project"`
	Type     string `json:"type"`
	Urlid    string `json:"urlid"`
	Testfile string `json:"testfile"`
	Status   string `json:"status"`
	Duration string `json:"duration"`
}

// type Test struct {
// 	ID       int    `json:"id"`
// 	Status   string `json:"status"`
// 	TestPath string `json:"testPath"`
// }