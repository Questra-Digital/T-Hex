package main

import (
	"cicd_runner/config"
	"cicd_runner/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Configure logging to output to stdout for Docker
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("Starting CICD Runner service...")

	config.LoadEnv()

	config.DBInit()

	router := gin.Default()

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	router.POST("/create_pipeline", routes.CreatePipeline)
	router.POST("/hook_callback", routes.CommitCallback)
	router.GET("/pipelines", routes.GetPipelines)

	log.Printf("CICD Runner service starting on port %s", config.EnvConfig.Port)
	router.Run(config.EnvConfig.Port)

}
