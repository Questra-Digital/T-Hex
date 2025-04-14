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
	errorFromGoogle := q.Get("error")

	// Check if there's an error parameter in the response
	if errorFromGoogle != "" {
		fmt.Printf("Error from Google: %s\n", errorFromGoogle)
		utils.RespondError(w, "Authorization failed: "+errorFromGoogle, http.StatusUnauthorized)
		return
	}

	// Validate state using the function from utils
	if !utils.ValidateState(stateFromGoogle) {
		fmt.Println("Error: Invalid state. Possible CSRF attack.")
		utils.RespondError(w, "Invalid state parameter. Possible CSRF attack.", http.StatusForbidden)
		return
	}

	// Check for a missing authorization code
	if codeFromGoogle == "" {
		fmt.Println("Error: Missing authorization code")
		utils.RespondError(w, "Invalid request: missing authorization code", http.StatusBadRequest)
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
		utils.RespondError(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	// Print tokens (for testing, store them securely in production)
	fmt.Fprintf(w, "Access Token: %s\n", tokens.AccessToken)

	http.Redirect(w, r, "/dashboard", http.StatusFound)
}
