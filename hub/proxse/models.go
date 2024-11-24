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
	time    int64
	method  string
	path    string
	reqBody string
	key     string
	proj    string
	status  int
	res     string
}

func (EventLogEntry) TableName() string {
	return "event_logs"
}
