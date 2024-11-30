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
	Key string `gorm:"primaryKey"`
}

func (ApiKey) TableName() string {
	return "api_keys"
}

// User
type User struct {
	Username string `gorm:"primaryKey"`
	Password string
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
type Event struct {
	Id        int64 `gorm:"primaryKey"`
	Time      int64
	SessionId string
	Method    string
	Path      string
	ReqBody   string
	Status    int
	Res       string
}

func (Event) TableName() string {
	return "events"
}

// Session ID to test Id mapping
type Session struct {
	SessionId string `gorm:"primaryKey"`
	TestId    int64
	Time      int64
	Valid     bool
	Status    bool
	Message   string
}

func (Session) TableName() string {
	return "sessions"
}

/// Session Id to Test Id mapping
type TestSession struct {
	TestId  int64 `gorm:"primaryKey;autoIncrement"`
	Time    int64
	Key     string
	Proj    string
	Current bool
}

func (TestSession) TableName() string {
	return "test_sessions"
}
