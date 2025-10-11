package routes

import (
	"bytes"
	"cicd_runner/config"
	"cicd_runner/db"
	"cicd_runner/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LastCommit struct {
	ID        string `json:"id"`        // SHA of commit
	Message   string `json:"message"`   // Commit message
	Timestamp string `json:"timestamp"` // ISO8601
}

type PushEventPayload struct {
	HeadCommit LastCommit `json:"head_commit"`
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
}

func CommitCallback(c *gin.Context) {
	var requestBody PushEventPayload

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		log.Println("Error binding JSON", err)
		return
	}

	pipelineId, err := db.GetPipelineIdByRepositoryPath(requestBody.Repository.FullName)
	if err != nil {
		log.Println("Error getting pipeline ID", err)
		return
	}

	pipelineEvent := models.PipelineEvent{
		PipelineID: pipelineId,
		Status:     "running",
		Timestamp:  time.Now(),
		Duration:   0,
		Details:    fmt.Sprintf("Commit to the %s branch", requestBody.Repository.FullName),
		Type:       "test",
	}

	err = db.CreatePipelineEvent(&pipelineEvent)
	if err != nil {
		log.Println("Error creating pipeline event", err)
		return
	}

	// Create the expected format for the frontend API
	frontendPayload := map[string]interface{}{
		"data": map[string]interface{}{
			"pipeline_id": pipelineId,
			"event":       pipelineEvent,
		},
	}

	jsonData, err := json.Marshal(frontendPayload)
	if err != nil {
		log.Printf("Error marshaling frontend payload: %v", err)
		return
	}

	// Send to the correct API endpoint
	frontendURL := config.EnvConfig.FrontendUrl + "/api/pipeline-status"
	log.Printf("Sending pipeline event to frontend: %s", frontendURL)
	log.Printf("Payload: %s", string(jsonData))

	resp, err := http.Post(frontendURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error sending pipeline event to frontend: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("Pipeline event sent to frontend: %s", resp.Status)
}
