package config

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"cicd_runner/models"
)

// Initial pipeline data (simplified for Go structs)
var initialPipelines = []struct {
	Pipeline models.Pipeline
	Events   []models.PipelineEvent
}{
	{
		Pipeline: models.Pipeline{
			Name:           "build-and-test",
			Description:    "Automated build and test pipeline for the main application",
			Status:         "running",
			LastRun:        parseTime("2024-01-15T10:30:00Z"),
			TriggerType:    "commit",
			BranchName:     "main",
			RepositoryPath: "github.com/company/frontend-app",
			Labels:         models.StringArray{"frontend", "testing", "automation"},
		},
		Events: []models.PipelineEvent{
			{
				Type:      "test",
				Status:    "success",
				Timestamp: parseTime("2024-01-15T10:30:00Z"),
				Duration:  int64(2*time.Minute + 34*time.Second),
				Details:   "Unit tests completed successfully with 0 errors",
			},
			{
				Type:      "test",
				Status:    "success",
				Timestamp: parseTime("2024-01-15T10:32:34Z"),
				Duration:  int64(1*time.Minute + 12*time.Second),
				Details:   "Integration tests passed - All 156 tests passed",
			},
			{
				Type:      "test",
				Status:    "running",
				Timestamp: parseTime("2024-01-15T10:33:46Z"),
				Duration:  int64(45 * time.Second),
				Details:   "E2E tests running...",
			},
		},
	},
	{
		Pipeline: models.Pipeline{
			Name:           "security-scan",
			Description:    "Security vulnerability scanning and code analysis",
			Status:         "success",
			LastRun:        parseTime("2024-01-15T09:15:00Z"),
			TriggerType:    "manual",
			BranchName:     "develop",
			RepositoryPath: "github.com/company/security-tools",
			Labels:         models.StringArray{"security", "analysis"},
		},
		Events: []models.PipelineEvent{
			{
				Type:      "test",
				Status:    "success",
				Timestamp: parseTime("2024-01-15T09:15:00Z"),
				Duration:  int64(5*time.Minute + 23*time.Second),
				Details:   "Security tests completed - no critical vulnerabilities found",
			},
		},
	},
	{
		Pipeline: models.Pipeline{
			Name:           "performance-test",
			Description:    "Performance testing and load testing pipeline",
			Status:         "failed",
			LastRun:        parseTime("2024-01-15T08:45:00Z"),
			TriggerType:    "commit",
			BranchName:     "feature/performance",
			RepositoryPath: "github.com/company/performance-suite",
			Labels:         models.StringArray{"performance", "testing"},
		},
		Events: []models.PipelineEvent{
			{
				Type:      "test",
				Status:    "failed",
				Timestamp: parseTime("2024-01-15T08:45:00Z"),
				Duration:  int64(3*time.Minute + 12*time.Second),
				Details:   "Performance tests failed - response time exceeded threshold",
			},
		},
	},
}

// Helper to parse ISO timestamps
func parseTime(ts string) time.Time {
	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		log.Fatalf("Invalid time format: %s", ts)
	}
	return t
}

var DB *gorm.DB

func DBInit() {
	dbUrl := EnvConfig.DbUrl

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	AutoMigrate(db)

	// Insert initial pipelines if DB is empty
	var count int64
	db.Model(&models.Pipeline{}).Count(&count)
	if count == 0 {
		log.Println("Inserting initial pipelines and events...")
		for _, item := range initialPipelines {
			// Create pipeline
			if err := db.Create(&item.Pipeline).Error; err != nil {
				log.Fatalf("Failed to insert pipeline: %s", err.Error())
			}

			// Create access token for the pipeline
			accessToken := models.AccessTokens{
				AccessToken: "initial-token-" + item.Pipeline.Name,
				PipelineID:  item.Pipeline.ID,
			}
			if err := db.Create(&accessToken).Error; err != nil {
				log.Fatalf("Failed to insert access token: %s", err.Error())
			}

			// Assign pipeline ID to events and insert them
			for i := range item.Events {
				item.Events[i].PipelineID = item.Pipeline.ID
				if err := db.Create(&item.Events[i]).Error; err != nil {
					log.Fatalf("Failed to insert pipeline event: %s", err.Error())
				}
			}
		}
		log.Println("Initial pipelines and events inserted successfully")
	}

	DB = db
	log.Println("Database initialization complete")
}

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(&models.Pipeline{}); err != nil {
		log.Fatalf("Failed to migrate Pipeline model: %s", err.Error())
	}
	if err := db.AutoMigrate(&models.PipelineEvent{}); err != nil {
		log.Fatalf("Failed to migrate PipelineEvent model: %s", err.Error())
	}
	if err := db.AutoMigrate(&models.AccessTokens{}); err != nil {
		log.Fatalf("Failed to migrate AccessTokens model: %s", err.Error())
	}
}
