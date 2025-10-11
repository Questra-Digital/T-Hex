package routes

import (
	"cicd_runner/config"
	"cicd_runner/db"
	"cicd_runner/github"
	"cicd_runner/models"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type PipelineRequest struct {
	Pipeline    models.Pipeline `json:"pipeline"`
	AccessToken string          `json:"access_token"`
}

func CreatePipeline(c *gin.Context) {

	var requestBody PipelineRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Printf("Failed to bind JSON request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	log.Printf("Creating pipeline: %s for repository: %s", requestBody.Pipeline.Name, requestBody.Pipeline.RepositoryPath)

	//No other input validation for now.
	pipeline := requestBody.Pipeline

	if err := db.CreatePipeline(&pipeline, requestBody.AccessToken); err != nil {
		log.Printf("Failed to create pipeline in database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	log.Printf("Pipeline created in database with ID: %d", pipeline.ID)

	repoPath := requestBody.Pipeline.RepositoryPath
	branchName := requestBody.Pipeline.BranchName
	accessToken := requestBody.AccessToken

	log.Printf("Creating webhook for repository: %s, branch: %s", repoPath, branchName)

	//No validations now for repoPath, branchName and accessToken
	owner, repo := extractOwnerAndRepo(repoPath)
	err := github.CreateWebhook(owner, repo, accessToken, config.EnvConfig.CallbackUrl, branchName, config.EnvConfig.Secret)
	if err != nil {
		log.Printf("Failed to create webhook: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	log.Printf("Pipeline created successfully with ID: %d", pipeline.ID)
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Pipeline created successfully", "data": map[string]uint{"pipeline_id": pipeline.ID}})
}

func extractOwnerAndRepo(repoPath string) (string, string) {
	parts := strings.Split(repoPath, "/")
	if len(parts) < 2 {
		return "", ""
	}
	return parts[0], parts[1]
}
