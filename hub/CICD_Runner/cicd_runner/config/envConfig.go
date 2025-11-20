package config

import (
	"log"
	"os"
	"sync"
)

var EnvConfig struct {
	Port            string
	DbUrl           string
	CallbackUrl     string
	Secret          string
	FrontendUrl     string
	ReverseProxyUrl string
}

// ProxySettings struct to store proxy configuration

type ProxySettings struct {
	TestId       string
	APIKey       string
	PipelineName string
	mu           sync.RWMutex
}

var ProxyConfig = &ProxySettings{}

func (c *ProxySettings) GetValues() (string, string, string) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.APIKey, c.PipelineName, c.TestId
}

func (c *ProxySettings) SetValues(apiKey, pipelineName, testId string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.APIKey = apiKey
	c.PipelineName = pipelineName
	c.TestId = testId
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

	EnvConfig.ReverseProxyUrl = os.Getenv("REVERSE_PROXY_URL")
	if EnvConfig.ReverseProxyUrl == "" {
		log.Fatalf("REVERSE_PROXY_URL is not set")
	}

	log.Printf("Env loaded successfully")

}
