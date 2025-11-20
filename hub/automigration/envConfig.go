package main

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

var EnvConfig struct {
	DbUrl string
}

func init() {
	godotenv.Load()
	EnvConfig.DbUrl = os.Getenv("DB_URL")
	if EnvConfig.DbUrl == "" {
		log.Fatalf("DB_URL is not set")
	}
	log.Printf("EnvConfig loaded successfully")
}