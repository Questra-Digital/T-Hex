// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type GetSettingInput struct {
	Browser               string `json:"browser"`
	Version               string `json:"version"`
	StepByStepDebugging   bool   `json:"stepByStepDebugging"`
	EnableLogs            bool   `json:"enableLogs"`
	Parallelism           bool   `json:"parallelism"`
	NumberOfParallelTests int    `json:"numberOfParallelTests"`
}
