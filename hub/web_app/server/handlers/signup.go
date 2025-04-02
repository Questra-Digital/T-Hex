package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"
	"server/utils"
)

// SignupRequest represents the expected JSON structure
type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	
	// Ensure request method is POST
	if r.Method != http.MethodPost {
		utils.RespondError(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var req SignupRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		utils.RespondError(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Validate input fields
	if req.Username == "" || req.Password == "" {
		utils.RespondError(w, "Username and Password required", http.StatusBadRequest)
		return
	}

	// Check if the user already exists
	if _, exists := models.ValidUsers[req.Username]; exists {
		utils.RespondError(w, "Username already taken", http.StatusConflict)
		return
	}

	// Save new user (for simplicity, adding to the in-memory map)
	models.ValidUsers[req.Username] = req.Password

	utils.RespondJSON(w, http.StatusCreated, map[string]string{"message": "User registered successfully"})
}
