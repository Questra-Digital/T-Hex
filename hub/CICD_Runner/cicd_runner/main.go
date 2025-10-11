package main

import (
	"cicd_runner/config"
	"cicd_runner/routes"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// Create HTTP server
	srv := &http.Server{
		Addr:    config.EnvConfig.Port,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("CICD Runner service starting on port %s", config.EnvConfig.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Give outstanding requests 30 seconds to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown HTTP server
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	// Close database connection
	if config.DB != nil {
		sqlDB, err := config.DB.DB()
		if err != nil {
			log.Printf("Error getting database instance: %v", err)
		} else {
			if err := sqlDB.Close(); err != nil {
				log.Printf("Error closing database connection: %v", err)
			} else {
				log.Println("Database connection closed successfully")
			}
		}
	}

	log.Println("Server exited gracefully")
}
