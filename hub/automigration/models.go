package main

import (
	"log"
	"time"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"strings"
	"fmt"
	"database/sql/driver"
)

// initializes DB
func DBInit() *gorm.DB {
	
	db, err := gorm.Open(postgres.Open(EnvConfig.DbUrl), &gorm.Config{
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

func DBClose(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to close database: %s", err.Error())
	}
	sqlDB.Close()
	log.Printf("Database closed successfully")
}

// User
type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

type User struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Role      UserRole  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Keys      *Keys     `gorm:"foreignKey:UserId;references:ID;constraint:OnDelete:CASCADE;" json:"keys"`
	Pipelines []Pipeline `gorm:"foreignKey:UserId;references:ID;constraint:OnDelete:CASCADE;" json:"pipelines"`
}

// TableName returns the database table name for the User model
func (User) TableName() string {
	return "users"
}

type Keys struct {
	ID     int64  `gorm:"primaryKey; autoIncrement" json:"id"`
	UserId int64  `gorm:"column:user_id;uniqueIndex;not null" json:"user_id"`
	APIKey string `gorm:"unique;not null" json:"api_key"`
}

func (Keys) TableName() string {
	return "user_keys"
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

type SessionSelenium struct {
	SessionId string `gorm:"primaryKey"`
	TestId    int64
	Time      int64
	Valid     bool
	Status    bool
	Message   string
	Events []Event `json:"events" gorm:"foreignKey:SessionId;references:SessionId;constraint:OnDelete:CASCADE;"`
}

func (SessionSelenium) TableName() string {
	return "sessions_selenium"
}

/// Session Id to Test Id mapping
type TestSession struct {
	TestId  int64 `gorm:"primaryKey;autoIncrement"`
	PipelineEventID int64 `gorm:"column:pipeline_event_id;not null" json:"pipeline_event_id"`
	Time    int64
	Key     string
	Proj    string
	Current bool
	SeleniumSessions []SessionSelenium `json:"selenium_sessions" gorm:"foreignKey:TestId;references:TestId;constraint:OnDelete:CASCADE;"`
}

func (TestSession) TableName() string {
	return "test_sessions"
}

// Pipelines///////////////////////////////////////////////////////

type StringArray []string

// Scan implements sql.Scanner interface (reading from DB)
func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = []string{}
		return nil
	}

	switch v := value.(type) {
	case string:
		v = strings.Trim(v, "{}")
		if v == "" {
			*a = []string{}
			return nil
		}
		*a = []string(strings.Split(v, ","))
		return nil
	case []byte:
		return a.Scan(string(v))
	default:
		return fmt.Errorf("cannot scan %T into StringArray", value)
	}
}

// Value implements driver.Valuer interface (writing to DB)
func (a StringArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "{}", nil
	}
	return "{" + strings.Join(a, ",") + "}", nil
}

type Pipeline struct {
	ID             uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	Status         string          `json:"status"`
	LastRun        time.Time       `json:"last_run"`
	TriggerType    string          `json:"trigger_type"`
	BranchName     string          `json:"branch_name"`
	RepositoryPath string          `json:"repository_path"`
	Labels         StringArray     `json:"labels" gorm:"type:text[]"`
	Events         []PipelineEvent `json:"events" gorm:"foreignKey:PipelineID;references:ID;constraint:OnDelete:CASCADE;"`
	AccessToken    *AccessTokens    `json:"access_token" gorm:"foreignKey:PipelineID;references:ID;constraint:OnDelete:CASCADE;"`
	UserId         int64            `json:"user_id" gorm:"column:user_id;not null"`
}

func (Pipeline) TableName() string {
	return "pipelines"
}

type PipelineEvent struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	PipelineID uint      `json:"pipeline_id"`
	Status     string    `json:"status"`
	Timestamp  time.Time `json:"timestamp"`
	Duration   int64     `json:"duration"` // in nanoseconds
	Details    string    `json:"details"`
	Type       string    `json:"type"`
	TestSession *TestSession `json:"test_session" gorm:"foreignKey:PipelineEventID;references:ID;constraint:OnDelete:CASCADE;"`
}

func (PipelineEvent) TableName() string {
	return "pipeline_events"
}

type AccessTokens struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	AccessToken string `json:"access_token"`
	PipelineID  uint   `json:"pipeline_id" gorm:"column:pipeline_id;not null;uniqueIndex"`
}

func (AccessTokens) TableName() string {
	return "access_tokens"
}