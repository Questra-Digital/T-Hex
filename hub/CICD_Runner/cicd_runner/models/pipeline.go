package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

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
	Events         []PipelineEvent `json:"events" gorm:"foreignKey:PipelineID;references:ID"`
	AccessToken    AccessTokens    `json:"access_token" gorm:"foreignKey:PipelineID;references:ID"`
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
}

func (PipelineEvent) TableName() string {
	return "pipeline_events"
}

type AccessTokens struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	AccessToken string `json:"access_token"`
	PipelineID  uint   `json:"pipeline_id"`
}

func (AccessTokens) TableName() string {
	return "access_tokens"
}
