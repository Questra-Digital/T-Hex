package config

import (
	"fmt"
	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from the .env file
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: No .env file found")
	}
}
