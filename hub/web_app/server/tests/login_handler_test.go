package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"server/handlers"
	"strings"
)

func TestLoginHandler_ValidLogin(t *testing.T) {

	reqBody := `{"username": "testuser", "password": "password123"}`
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handlers.LoginHandler(rec, req)

	// Assert status code
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	// Parse the response body
	var resp map[string]string
	err := json.NewDecoder(rec.Body).Decode(&resp)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// Assert the "message" field in the response
	if resp["message"] != "Login successful" {
		t.Errorf("Expected message 'Login successful', got '%s'", resp["message"])
	}

	// Check for the Set-Cookie header
	cookie := rec.Header().Get("Set-Cookie")
	if cookie == "" {
		t.Errorf("Expected Set-Cookie header in response, but got none")
	} else {
	
		// Verify that the cookie contains the "token" field and it's not empty
		if !containsToken(cookie) {
			t.Errorf("Expected cookie to contain a non-empty 'token' field, but it was missing or empty")
		}
	}
}

func containsToken(cookie string) bool {

	// Parse the cookie string into http.Cookie objects
	cookies := strings.Split(cookie, ";")
	for _, c := range cookies {
		// Parse key-value pairs
		parts := strings.SplitN(strings.TrimSpace(c), "=", 2)
		if len(parts) == 2 && parts[0] == "token" && parts[1] != "" {
			return true
		}
	}
	return false
}

func TestLoginHandler_InvalidUsername(t *testing.T) {

	reqBody := `{"username": "wronguser", "password": "password123"}`
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handlers.LoginHandler(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", rec.Code)
	}
}

func TestLoginHandler_InvalidPassword(t *testing.T) {

	reqBody := `{"username": "testuser", "password": "wrongpassword"}`
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handlers.LoginHandler(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", rec.Code)
	}
}

func TestLoginHandler_EmptyRequestBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handlers.LoginHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", rec.Code)
	}
}
