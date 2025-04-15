package handlers

import (
	"encoding/json"
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

		// Delete the state before redirecting to GoogleLogin
		if stateFromGoogle != "" {
			utils.DeleteState(stateFromGoogle)
		}

		// Redirect to Google Login to restart the OAuth process
		http.Redirect(w, r, "/googleLogin", http.StatusFound)
		return
	}

	// Validate state using the function from utils
	if !utils.ValidateState(stateFromGoogle) {
		fmt.Println("Error: Invalid state. Possible CSRF attack.")

		// Redirect to Google Login to restart the OAuth process
		http.Redirect(w, r, "/googleLogin", http.StatusFound)
		return
	}

	// Check for a missing authorization code
	if codeFromGoogle == "" {
		fmt.Println("Error: Missing authorization code")

		// Delete the state before redirecting to GoogleLogin
		if stateFromGoogle != "" {
			utils.DeleteState(stateFromGoogle)
		}

		// Redirect to Google Login to restart the OAuth process
		http.Redirect(w, r, "/googleLogin", http.StatusFound)
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

		// Delete the state before redirecting to GoogleLogin
		utils.DeleteState(stateFromGoogle)

		// Redirect to Google Login to restart the OAuth process
		http.Redirect(w, r, "/googleLogin", http.StatusFound)
		return
	}

	// Print tokens (for testing, store them securely in production)
	fmt.Printf("Access Token: %s\n", tokens.AccessToken)
	fmt.Printf("Refresh Token: %s\n", tokens.RefreshToken)

	oauth2Client := oauth2Config.Client(r.Context(),tokens)

	userInfoEndpoint := "https://www.googleapis.com/oauth2/v2/userinfo"

	resp,err := oauth2Client.Get(userInfoEndpoint)

	if err!= nil{
		fmt.Printf("Error fetching user info: %v\n",err)
		// Delete the state before redirecting to GoogleLogin
		if stateFromGoogle != "" {
			utils.DeleteState(stateFromGoogle)
		}
		http.Redirect(w,r,"/googleLogin",http.StatusFound)
		return
	}

	//To deallocate resource for opening json body for an http connection
	defer resp.Body.Close()

	userInfo := struct{
		Name string `json:"name"`
		Email string `json:"email"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	
	if err !=nil{
		fmt.Printf("Error decoding user info: %v\n",err)
		http.Redirect(w,r,"/googleLogin",http.StatusFound)
		return
	}

	fmt.Printf("Welcome, %s (%s)\n", userInfo.Name, userInfo.Email)
}
