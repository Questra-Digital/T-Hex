package model

type Test struct {
	ID       int    `json:"id"`
	Status   string `json:"status"`
	TestPath string `json:"testPath"`
}