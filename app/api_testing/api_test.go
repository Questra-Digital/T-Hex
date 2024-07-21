package api_testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

// Define a struct to represent the response structure
type ApiResponse struct {
	Message string `json:"message"`
}

func TestCreateProjectInfoAPI(t *testing.T) {
	// Define the request body
	requestBody := map[string]interface{}{
		"GitEmail":            "test@example.com",
		"GitProjectName":      "TestProject",
		"GitLanguage":         "Go",
		"GitNoOfTestCases":    10,
		"GitTestCaseFileName": "test_cases.txt",
	}

	// Marshal the request body to JSON
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "http://localhost:9090", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Unmarshal the response body
	var apiResponse ApiResponse
	if err := json.Unmarshal(respBody, &apiResponse); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	// Check the response message
	expectedMessage := "Project info created successfully"
	if apiResponse.Message != expectedMessage {
		t.Errorf("Expected message %q, got %q", expectedMessage, apiResponse.Message)
	}
}

func TestMain(m *testing.M) {
	// Run the tests
	fmt.Println("Running API tests...")
	testResult := m.Run()
	os.Exit(testResult)
}
