package routes

import (
	"cicd_runner/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FetchTestSessionsRequest struct {
	PipelineId      string `json:"pipeline_id"`
	PipelineEventId string `json:"pipeline_event_id"`
	UserId          int64  `json:"user_id"`
}

func FetchTestSession(c *gin.Context) {
	var request FetchTestSessionsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Parse string IDs to uint
	pipelineId, err := strconv.ParseUint(request.PipelineId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid pipeline_id"})
		return
	}

	pipelineEventId, err := strconv.ParseUint(request.PipelineEventId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid pipeline_event_id"})
		return
	}

	userId, err := db.VerifyUserId(request.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	testSession, err := db.GetTestSession(uint(pipelineId), uint(pipelineEventId), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Test session fetched successfully", "data": gin.H{"test_session": testSession}})
}
