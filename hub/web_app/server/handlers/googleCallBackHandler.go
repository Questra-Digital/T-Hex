package handlers

import (
	"fmt"
	"net/http"
	"os"
	"server/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

)


// GoogleCallback handles the OAuth callback after Google authorization
func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters from Google's callback
	q := r.URL.Query()
	stateFromGoogle := q.Get("state")
	codeFromGoogle := q.Get("code")

	// Validate state using the function from utils
	if !utils.ValidateState(stateFromGoogle) {
		fmt.Println("Error: Invalid state. Possible CSRF attack.")
		return
	}

	// OAuth2 configuration
	oauth2Config := oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	// Exchange authorization code for access token
	tokens, err := oauth2Config.Exchange(r.Context(), codeFromGoogle)
	if err != nil {
		fmt.Println("Error: Failed to exchange token")
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	// Print tokens (for testing, store them securely in production)
	fmt.Fprintf(w, "Access Token: %s\n", tokens.AccessToken)
}
