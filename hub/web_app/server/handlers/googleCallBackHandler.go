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

func ErrorRedirectToLogin(errorBody string, stateFromGoogle string,w http.ResponseWriter,r *http.Request){
	fmt.Printf("%s",errorBody)
	// Delete the state before redirecting to GoogleLogin
	if stateFromGoogle != "" {
		utils.DeleteState(stateFromGoogle)
	}
	// Redirect to Google Login to restart the OAuth process
	http.Redirect(w, r, "/googleLogin", http.StatusFound)
}

// GoogleCallback handles the OAuth callback after Google authorization
func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters from Google's callback
	q := r.URL.Query()
	stateFromGoogle := q.Get("state")
	codeFromGoogle := q.Get("code")
	errorFromGoogle := q.Get("error")

	// Check if there's an error parameter in the response
	if errorFromGoogle != "" {
		errorBody := fmt.Sprintf("Error from Google: %s\n", errorFromGoogle)
		ErrorRedirectToLogin(errorBody, stateFromGoogle, w, r)
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
		errorBody := "Error: Missing authorization code"
		ErrorRedirectToLogin(errorBody, stateFromGoogle, w, r)
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
		errorBody := "Error: Failed to exchange token"
		ErrorRedirectToLogin(errorBody, stateFromGoogle, w, r)
		return
	}

	// Print tokens (for testing, store them securely in production)
	fmt.Printf("Access Token: %s\n", tokens.AccessToken)
	fmt.Printf("Refresh Token: %s\n", tokens.RefreshToken)

	oauth2Client := oauth2Config.Client(r.Context(),tokens)

	userInfoEndpoint := "https://www.googleapis.com/oauth2/v2/userinfo"

	resp,err := oauth2Client.Get(userInfoEndpoint)

	if err!= nil{
		errorBody := fmt.Sprintf("Error fetching user info: %v\n",err)
		ErrorRedirectToLogin(errorBody, stateFromGoogle, w, r)
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
		errorBody := fmt.Sprintf("Error decoding user info: %v\n",err)
		ErrorRedirectToLogin(errorBody, stateFromGoogle, w, r)
		return
	}

	fmt.Printf("Welcome, %s (%s)\n", userInfo.Name, userInfo.Email)
}
