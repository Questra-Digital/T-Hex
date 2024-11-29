package main

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// initializes DB
func DBInit(connectionStr string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}
	return db
}

type Tabler interface {
	TableName() string
}

// Row from api_keys database table. Only currently valid Key will exist
type ApiKey struct {
	gorm.Model
	Key string
}

func (ApiKey) TableName() string {
	return "api_keys"
}

// User
type User struct {
	Username string `gorm:"primaryKey;not null"`
	Password string `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}

// User to API Key relations. Invalidated keys will also exist here
type UserKey struct {
	Username string `gorm:"primaryKey"`
	Key      string `gorm:"primaryKey"`
}

func (UserKey) TableName() string {
	return "users_keys"
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

// Session ID to Key mapping
type KeySession struct {
	gorm.Model
	Time      int64
	Key       string
	Proj      string
	SessionId string
	Valid     bool
}

func (KeySession) TableName() string {
	return "key_sessions"
}
