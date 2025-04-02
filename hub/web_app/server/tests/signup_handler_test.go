package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"server/handlers"
	"reflect"
)

func testSignupHandler(t *testing.T, reqBody map[string]string, expectedStatus int, expectedBody map[string]string) {
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handlers.SignupHandler(rec, req)

	if rec.Code != expectedStatus {
		t.Errorf("Expected status %d, got %d", expectedStatus, rec.Code)
	}

	var actualBody map[string]string
	json.Unmarshal(rec.Body.Bytes(), &actualBody)

	// Compare the two JSON maps
	if !reflect.DeepEqual(actualBody, expectedBody) {
		t.Errorf("Expected response body '%v', got '%v'", expectedBody, actualBody)
	}
}


func TestSignupHandler_Success(t *testing.T) {
	requestBody := map[string]string{
		"username": "newuser",
		"password": "securepassword",
	}
	expectedBody := map[string]string{
		"message": "User registered successfully",
	}

	testSignupHandler(t, requestBody, http.StatusCreated, expectedBody)
}

func TestSignupHandler_UserAlreadyExists(t *testing.T) {
	requestBody := map[string]string{
		"username": "testuser", // Already exists in models.ValidUsers
		"password": "password123",
	}
	expectedBody := map[string]string{
		"error": "Username already taken",
	}

	testSignupHandler(t, requestBody, http.StatusConflict, expectedBody)
}

func TestSignupHandler_MissingFields(t *testing.T) {
	requestBody := map[string]string{
		"username": "", 
		"password": "",
	}
	expectedBody := map[string]string{
		"error": "Username and Password required",
	}

	testSignupHandler(t, requestBody, http.StatusBadRequest, expectedBody)
}

func TestSignupHandler_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer([]byte("{invalid_json}")))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handlers.SignupHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", rec.Code)
	}

	expectedBody := map[string]string{"error": "Invalid JSON format"}
	var actualBody map[string]string
	_ = json.Unmarshal(rec.Body.Bytes(), &actualBody)

	if actualBody["error"] != expectedBody["error"] {
		t.Errorf("Expected response body %v, got %v", expectedBody, actualBody)
	}
}
