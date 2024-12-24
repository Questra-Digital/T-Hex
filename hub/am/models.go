package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// initializes DB
func DBInit() *gorm.DB {
	dbStr := os.Getenv("DB")
	if dbStr == "" {
		dbStr = "postgres://thex:thex1234@db/thex"
	}
	db, err := gorm.Open(postgres.Open(dbStr), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)
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

// ContactUs message
type ContactUsMessage struct {
	Id      int64 `gorm:"primaryKey;autoIncrement"`
	Email   string
	Message string
}

func (ContactUsMessage) TableName() string {
	return "contactus_messages"
}
