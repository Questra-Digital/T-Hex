package config

import (
	"os"
	"github.com/joho/godotenv"
	"strconv"
	"log"
	"fmt"
)

var EnvConfig struct {
	SeleniumHubUrl string
	Port           string
	CacheTTL       int32
	DbUrl          string
}

func EnvConfigInit() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: No .env file found, using environment variables")
	}

	//Load environment variables into EnvConfig
	EnvConfig.SeleniumHubUrl = os.Getenv("SELENIUM_HUB_URL")
	if EnvConfig.SeleniumHubUrl == "" {
		log.Fatalf("SELENIUM_HUB_URL is not set")
	}
	EnvConfig.Port = os.Getenv("PORT")
	if EnvConfig.Port == "" {
		log.Fatalf("PORT is not set")
	}
	//Get CACHE_TTL from environment variable into int32
	ttlStr := os.Getenv("CACHE_TTL")
	if ttlStr == "" {
		log.Fatalf("CACHE_TTL is not set")
	}
	i, err := strconv.Atoi(ttlStr)
	if err != nil {
		log.Fatalf("Failed to parse CACHE_TTL: %s", err.Error())
	}
	EnvConfig.CacheTTL = int32(i)

	EnvConfig.DbUrl = os.Getenv("DB_URL")
	if EnvConfig.DbUrl == "" {
		log.Fatalf("DB_URL is not set")
	}
	log.Printf("EnvConfig loaded successfully")
}
