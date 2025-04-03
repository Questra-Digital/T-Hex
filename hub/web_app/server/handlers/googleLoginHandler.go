package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"server/utils"
	"time"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"server/models"
)


// googleLoginHandler redirects users to Google's OAuth 2.0 consent page.
	func GoogleLogin(w http.ResponseWriter, r *http.Request) {
		
		// Ensure CLIENT_ID and CLIENT_SECRET are set
		clientID := os.Getenv("CLIENT_ID")
		clientSecret := os.Getenv("CLIENT_SECRET")
		if clientID == "" || clientSecret == "" {
			fmt.Println("Missing CLIENT_ID or CLIENT_SECRET")
			utils.RespondError(w, "Missing CLIENT_ID or CLIENT_SECRET", http.StatusInternalServerError)
			return
		}

		oauth2Config := oauth2.Config{
			ClientID:      os.Getenv("CLIENT_ID"),
			ClientSecret:  os.Getenv("CLIENT_SECRET"),
			RedirectURL:  "http://localhost:8080/auth/google/callback",
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.profile",
				"https://www.googleapis.com/auth/userinfo.email",
			},
			Endpoint: google.Endpoint,
		}

		// Generate a random state for CSRF protection
		src := rand.NewSource(time.Now().UnixNano()) // Create a new random source
		rng := rand.New(src)                         // Create a new random generator
		state := fmt.Sprintf("%d", rng.Int63())      // Generate a random state string

		// Store state in the in-memory map
		models.StateStore.Lock()
		models.StateStore.Map[state] = state
		models.StateStore.Unlock()

		// Generate OAuth URL
		authURL := oauth2Config.AuthCodeURL(state, oauth2.AccessTypeOffline)

		// Redirect the user to Google OAuth consent page
		http.Redirect(w, r, authURL, http.StatusFound)
}