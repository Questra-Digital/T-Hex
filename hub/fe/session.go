package main

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"sync"
	"time"
)

type SessionData struct {
	Username  string
	ExpiresAt time.Time
}

var tokenStore = make(map[string]SessionData)
var tokenStoreMu sync.RWMutex

func SessionCreate(username string) string {
	for {
		tokenBytes := make([]byte, 64)
		_, err := rand.Read(tokenBytes)
		if err != nil {
			panic("Error generating random token:")
		}
		token := hex.EncodeToString(tokenBytes)
		tokenStoreMu.RLock()
		if _, exists := tokenStore[token]; exists {
			tokenStoreMu.RUnlock()
			continue
		}
		tokenStoreMu.RUnlock()
		expiresAt := time.Now().Add(24 * time.Hour)
		session := SessionData{
			Username:  username,
			ExpiresAt: expiresAt,
		}
		tokenStoreMu.Lock()
		tokenStore[token] = session
		tokenStoreMu.Unlock()
		return token
	}
}

func SessionInvalidate(token string) {
	tokenStoreMu.Lock()
	delete(tokenStore, token)
	tokenStoreMu.Unlock()
}

func SessionAutoInvalidate() {
	for token, session := range tokenStore {
		if time.Now().After(session.ExpiresAt) {
			SessionInvalidate(token)
		}
	}
}

func startSessionAutoInvalidation() {
	ticker := time.NewTicker(30 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		SessionAutoInvalidate()
	}
}

// Set a secure session for a user
func SetSessionUser(w http.ResponseWriter, username string) {
	token := SessionCreate(username)
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	})
}

// Get the logged-in user from the session, or empty string
func GetSessionUser(r *http.Request) string {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return ""
	}
	token := cookie.Value
	tokenStoreMu.Lock()
	session, exists := tokenStore[token]
	tokenStoreMu.Unlock()
	if !exists {
		return ""
	}
	return session.Username
}

// Clear the session for a user
func ClearSessionUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err == nil {
		tokenStoreMu.Lock()
		delete(tokenStore, cookie.Value)
		tokenStoreMu.Unlock()
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
