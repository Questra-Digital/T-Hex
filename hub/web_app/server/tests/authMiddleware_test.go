package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"server/handlers"
	"server/middleware"
	"server/utils"
	"testing"
)

// Helper function to test AuthMiddleware
func testAuthMiddleware(t *testing.T, token string, expectedStatus int, expectedBody map[string]string) {
	req := httptest.NewRequest(http.MethodGet, "/dashboard", nil)
	if token != "" {
		req.Header.Set("Cookie", "token="+token)
	}

	rec := httptest.NewRecorder()

	handler := middleware.AuthMiddleware(handlers.DashboardHandler)

	handler.ServeHTTP(rec, req)

	// Assert the response status code
	if rec.Code != expectedStatus {
		t.Errorf("Expected status %d, got %d", expectedStatus, rec.Code)
	}

	// Parse the actual response body into a map
	var actualBody map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &actualBody)
	if err != nil {
		t.Fatalf("Failed to parse actual response body: %v", err)
	}

	// Compare the two JSON maps
	if !reflect.DeepEqual(actualBody, expectedBody) {
		t.Errorf("Expected response body '%v', got '%v'", expectedBody, actualBody)
	}
}

func TestAuthMiddleware_ValidToken(t *testing.T) {
	token, _ := utils.GenerateJWT("testuser")
	expectedBody := map[string]string{
		"message": "Welcome to the Dashboard, testuser!",
	}
	testAuthMiddleware(t, token, http.StatusOK, expectedBody)
}

func TestAuthMiddleware_MissingToken(t *testing.T) {
	expectedBody := map[string]string{
		"error": "Missing Token",
	}
	testAuthMiddleware(t, "", http.StatusUnauthorized, expectedBody)
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	token := "invalidToken"
	expectedBody := map[string]string{
		"error": "Invalid Token",
	}
	testAuthMiddleware(t, token, http.StatusUnauthorized, expectedBody)
}

func TestAuthMiddleware_NoUsernameInToken(t *testing.T) {
	token, _ := utils.GenerateJWT("") // Generate a token with no username
	expectedBody := map[string]string{
		"error": "Missing Username",
	}
	testAuthMiddleware(t, token, http.StatusUnauthorized, expectedBody)
}

func TestAuthMiddleware_InvalidUsername(t *testing.T) {
	// Generate a valid JWT with a username that does NOT exist in ValidUsers
	token, _ := utils.GenerateJWT("nonexistentuser") 

	expectedBody := map[string]string{
		"error":"User does not exist",
	}
	
	testAuthMiddleware(t, token, http.StatusUnauthorized, expectedBody)
}