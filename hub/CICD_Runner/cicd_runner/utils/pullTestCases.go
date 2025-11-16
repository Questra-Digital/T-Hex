package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// GitHubFile represents a file in GitHub repository
type GithubFile struct {
	Name        string `json:"name"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
}

// GitHubAPIResponse represents the response from GitHub API
type GitHubAPIResponse []GithubFile

// PullTestCases fetches test cases from GitHub repository and stores them locally
func PullTestCases(repoPath string, accessToken string) error {
	// Parse repository path (format: owner/repo)
	parts := strings.Split(repoPath, "/")
	if len(parts) != 2 {
		return fmt.Errorf("invalid repository path format. Expected: owner/repo")
	}

	owner := parts[0]
	repo := parts[1]

	// Check if test_cases directory exists if it doesn't exist, throw error
	testCasesDir := "/app/test_cases"
	if _, err := os.Stat(testCasesDir); os.IsNotExist(err) {
		return fmt.Errorf("test_cases directory not found")
	}

	// Check if /test directory exists in the repository
	testFiles, err := fetchTestDirectoryContents(owner, repo, accessToken)
	if err != nil {
		return fmt.Errorf("failed to fetch test directory contents: %v", err)
	}

	if len(testFiles) == 0 {
		return fmt.Errorf("no files found in test directory")
	}

	// Validate that selenium test files are present
	seleniumTestsFound := false
	for _, file := range testFiles {
		if isSeleniumTestFile(file.Name) {
			seleniumTestsFound = true
			break
		}
	}

	if !seleniumTestsFound {
		return fmt.Errorf("no selenium test files found in /test directory")
	}

	//Download the test files to the test_cases directory
	downloadedFiles := []string{}
	for _, file := range testFiles {
		filePath, err := downloadFile(file.DownloadURL, file.Name, testCasesDir)
		if err != nil {
			return fmt.Errorf("failed to download file: %v", err)
		}
		downloadedFiles = append(downloadedFiles, filePath)
	}

	log.Printf("Successfully downloaded %d test files to %s", len(downloadedFiles), testCasesDir)

	return nil
}

func downloadFile(downloadURL, fileName, testCasesDir string) (string, error) {
	resp, err := http.Get(downloadURL)
	if err != nil {
		return "", fmt.Errorf("failed to download file: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	filePath := filepath.Join(testCasesDir, fileName)
	
	err = os.WriteFile(filePath, body, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to write file: %v", err)
	}

	log.Printf("Successfully downloaded file: %s to %s", fileName, filePath)
	return filePath, nil
}

// checkTestDirectoryExists checks if /test directory exists in the repository
func fetchTestDirectoryContents(owner, repo, accessToken string) ([]GithubFile, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/tests", owner, repo)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "token "+accessToken)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("test directory not found in repository")
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	var githubAPIResponse GitHubAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&githubAPIResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return githubAPIResponse, nil
}

// isSeleniumTestFile checks if a file is a selenium test file
func isSeleniumTestFile(filename string) bool {
	return strings.Contains(strings.ToLower(filename), "test")
}
