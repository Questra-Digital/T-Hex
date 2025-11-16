package utils

import (
	"bytes"
	"cicd_runner/config"
	"cicd_runner/models"
	"encoding/json"
	"log"
	"net/http"
)

func SendPipelineEvent(pipelineEvent models.PipelineEvent) error {
	// Create the expected format for the frontend API
	frontendPayload := map[string]interface{}{
		"data": map[string]interface{}{
			"pipeline_id": pipelineEvent.PipelineID,
			"event":       pipelineEvent,
		},
	}
	jsonData, err := json.Marshal(frontendPayload)
	if err != nil {
		log.Printf("Error marshaling frontend payload: %v", err)
		return err
	}

	// Send to the correct API endpoint
	frontendURL := config.EnvConfig.FrontendUrl + "/api/pipeline-status"
	log.Printf("Sending pipeline event to frontend: %s", frontendURL)
	log.Printf("Payload: %s", string(jsonData))

	resp, err := http.Post(frontendURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error sending pipeline event to frontend: %v", err)
		return err
	}
	defer resp.Body.Close()

	log.Printf("Pipeline event sent to frontend: %s", resp.Status)

	return nil
}
