package main

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"sync"
)

var sessionStore = struct {
	sync.Mutex
	data map[string]string
}{data: make(map[string]string)}

// Generate a secure random token
func generateToken() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		panic("Unable to generate secure token")
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

// Set a secure session for a user
func SetSessionUser(w http.ResponseWriter, username string) {
	token := generateToken()

	// Save the username in the server-side session store
	sessionStore.Lock()
	sessionStore.data[token] = username
	sessionStore.Unlock()

	// Set a secure, HTTP-only cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true, // Use Secure only with HTTPS
	})
}

// Get the logged-in user from the session
func GetSessionUser(r *http.Request) string {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return ""
	}

	token := cookie.Value

	// Retrieve the username from the session store
	sessionStore.Lock()
	username, exists := sessionStore.data[token]
	sessionStore.Unlock()

	if !exists {
		return ""
	}
	return username
}

// Clear the session for a user
func ClearSessionUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err == nil {
		// Remove the session from the store
		sessionStore.Lock()
		delete(sessionStore.data, cookie.Value)
		sessionStore.Unlock()
	}

	// Invalidate the cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		//Secure:   true, // Use Secure only with HTTPS
	})
}
