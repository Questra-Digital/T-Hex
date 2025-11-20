package routes

import (
	"cicd_runner/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetPipelinesRequest struct {
	UserId int64 `json:"user_id"`
}

func GetPipelines(c *gin.Context) {

	var requestBody GetPipelinesRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	userId, err := db.VerifyUserId(requestBody.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	pipelines, err := db.GetPipelines(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Pipelines fetched successfully", "data": gin.H{"pipelines": pipelines}})
}
