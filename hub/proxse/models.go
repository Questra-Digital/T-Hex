package main

import (
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

// Row from api_keys database table
type ApiKey struct {
	gorm.Model
	Key string
}

func (ApiKey) TableName() string {
	return "api_keys"
}

// Event log entry
type EventLogEntry struct {
	gorm.Model
	Time    int64
	Method  string
	Path    string
	ReqBody string
	Key     string
	Proj    string
	Status  int
	Res     string
}

func (EventLogEntry) TableName() string {
	return "event_logs"
}
