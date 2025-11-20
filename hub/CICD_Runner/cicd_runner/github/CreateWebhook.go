package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type WebhookRequest struct {
	Name   string   `json:"name"`
	Active bool     `json:"active"`
	Events []string `json:"events"`
	Config struct {
		URL         string `json:"url"`
		ContentType string `json:"content_type"`
		Secret      string `json:"secret,omitempty"`
	} `json:"config"`
}

type GitHubError struct {
	Message string `json:"message"`
}

func CreateWebhook(owner, repo, accessToken, callbackURL, branch, secret string) error {

	// First, validate that the repository exists and the token has access
	if err := validateRepositoryAccess(owner, repo, accessToken); err != nil {
		log.Printf("Repository validation failed for %s/%s: %v", owner, repo, err)
		return fmt.Errorf("repository validation failed: %v", err)
	}

	reqBody := WebhookRequest{
		Name:   "web",
		Active: true,
		Events: []string{"push"},
	}
	reqBody.Config.URL = callbackURL
	reqBody.Config.ContentType = "json"
	reqBody.Config.Secret = secret

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal webhook request: %v", err)
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/hooks", owner, repo)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "token "+accessToken)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request to GitHub API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 201 {
		log.Printf("Webhook created successfully for %s/%s, status: %s", owner, repo, resp.Status)
		return nil
	}

	// Read the response body for detailed error information
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body for webhook creation: %v", err)
		return fmt.Errorf("failed to create webhook, status: %s (could not read error details)", resp.Status)
	}

	log.Printf("Webhook creation failed with status %d, response body: %s", resp.StatusCode, string(body))

	var githubError GitHubError
	if err := json.Unmarshal(body, &githubError); err == nil && githubError.Message != "" {
		log.Printf("GitHub API error details: %s", githubError.Message)
		return fmt.Errorf("failed to create webhook: %s", githubError.Message)
	}

	return fmt.Errorf("failed to create webhook, status: %s, response: %s", resp.Status, string(body))
}

func validateRepositoryAccess(owner, repo, accessToken string) error {

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Failed to create validation request for %s/%s: %v", owner, repo, err)
		return fmt.Errorf("failed to create validation request: %v", err)
	}

	req.Header.Set("Authorization", "token "+accessToken)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to make validation request for %s/%s: %v", owner, repo, err)
		return fmt.Errorf("failed to validate repository access: %v", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		log.Printf("Repository %s/%s validation successful", owner, repo)
		return nil // Repository exists and token has access
	case 401:
		log.Printf("Invalid or expired token for repository %s/%s", owner, repo)
		return fmt.Errorf("invalid or expired GitHub token")
	case 403:
		log.Printf("Token does not have access to repository %s/%s", owner, repo)
		return fmt.Errorf("token does not have access to repository %s/%s", owner, repo)
	case 404:
		log.Printf("Repository %s/%s not found or no access", owner, repo)
		return fmt.Errorf("repository %s/%s not found or token does not have access", owner, repo)
	default:
		log.Printf("Unexpected status code %d when validating repository %s/%s", resp.StatusCode, owner, repo)
		return fmt.Errorf("unexpected status code %d when validating repository", resp.StatusCode)
	}
}
