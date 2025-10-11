package routes

import (
	"cicd_runner/db"
	"cicd_runner/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetPipelines(c *gin.Context) {

	pipelines, err := db.GetPipelines()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Pipelines fetched successfully", "data": map[string][]models.Pipeline{"pipelines": pipelines}})
}