package routes

import (
	"cicd_runner/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FetchSeleniumEventsRequest struct {
	PipelineId      string `json:"pipeline_id"`
	PipelineEventId string `json:"pipeline_event_id"`
	UserId          int64  `json:"user_id"`
	SessionId       string `json:"session_id"`
}

func FetchSeleniumEvents(c *gin.Context) {
	var request FetchSeleniumEventsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	//1. Verify the user id
	userId, err := db.VerifyUserId(request.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	//1. Fetch the test session
	pipelineId, err := strconv.ParseUint(request.PipelineId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid pipeline_id"})
		return
	}
	pipelineEventId, err := strconv.ParseUint(request.PipelineEventId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid pipeline_event_id"})
		return
	}

	//2. Fetch the test session
	testSession, err := db.GetTestSession(uint(pipelineId), uint(pipelineEventId), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	//3. Fetch the selenium session by session id
	seleniumSession, err := db.GetSeleniumSessionBySessionId(request.SessionId, testSession.TestId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	//4. Fetch the selenium events for the session
	seleniumEvents, err := db.GetSeleniumEventsBySessionId(seleniumSession.SessionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Selenium events fetched successfully", "data": gin.H{"selenium_events": seleniumEvents}})
}
