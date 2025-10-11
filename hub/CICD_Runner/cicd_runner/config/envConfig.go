package config

import (
	"os"
)

import "log"

var EnvConfig struct {
	Port        string
	DbUrl       string
	CallbackUrl string
	Secret      string
	FrontendUrl string
}

func LoadEnv() {

	EnvConfig.Port = os.Getenv("PORT")
	if EnvConfig.Port == "" {
		log.Fatalf("PORT is not set")
	}
	EnvConfig.DbUrl = os.Getenv("DB_URL")
	if EnvConfig.DbUrl == "" {
		log.Fatalf("DB_URL is not set")
	}
	EnvConfig.CallbackUrl = os.Getenv("CALLBACK_URL")
	if EnvConfig.CallbackUrl == "" {
		log.Fatalf("CALLBACK_URL is not set")
	}
	EnvConfig.Secret = os.Getenv("SECRET")
	if EnvConfig.Secret == "" {
		log.Fatalf("SECRET is not set")
	}

	EnvConfig.FrontendUrl = os.Getenv("FRONTEND_URL")
	if EnvConfig.FrontendUrl == "" {
		log.Fatalf("FRONTEND_URL is not set")
	}

	log.Printf("Env loaded successfully")

}
