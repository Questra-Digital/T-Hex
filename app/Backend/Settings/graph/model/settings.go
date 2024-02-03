package model

type Settings struct {
	ID                  int    `json:"id"`
	Browser             string `json:"browser"`
	Version             string `json:"version"`
	StepByStepDebugging bool   `json:"stepByStepDebugging"`
	EnableLogs          bool   `json:"enableLogs"`
	Parallelism         bool   `json:"parallelism"`
}
