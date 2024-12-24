package main

import (
	"crypto/rand"
	"encoding/hex"
	"net/mail"
	"sync"
	"time"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// hashes the password using bcrypt
func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),
		bcrypt.DefaultCost)
	return string(bytes), err
}

// compares a hashed password with a plain-text password
func PasswordHashMatch(hash, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	return err == nil
}

// whether a string is a valid email address
func EmailIsValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func UsernameIsValid(username string) bool {
	if len(username) < 4 {
		return false
	}
	for i, r := range username {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			continue
		}
		if r == '_' && i > 0 && i < len(username)-1 {
			continue
		}
		return false
	}
	return true
}

func UsernameExists(username string) bool {
	var user User
	if err := db.First(&user, "username = ?", username).Error; err != nil {
		return false
	}
	return true
}

func UserGetCurrentKey(username string) string {
	var result struct {
		Key string
	}
	err := db.Table("users_keys").
		Select("api_keys.key").
		Joins("JOIN api_keys ON users_keys.key = api_keys.key").
		Where("users_keys.username = ?", username).
		Scan(&result).Error
	if err != nil {
		return ""
	}
	return result.Key
}

var SingleUseTokens = struct {
	sync.RWMutex
	data map[string]struct{Time time.Time; Key string}
}{data: make(map[string]struct{Time time.Time; Key string})}

func SingleUseTokenGenerate(username string) string {
	if data, exists := SingleUseTokens.data[username]; exists &&
			time.Now().After(data.Time) {
		return "" // refuse to generate new if previous not expired
	}
	for {
		tokenBytes := make([]byte, 64)
		_, err := rand.Read(tokenBytes)
		if err != nil {
			panic("Error generating random token:" + err.Error())
		}
		token := hex.EncodeToString(tokenBytes)
		SingleUseTokens.RLock()
		if _, exists := SingleUseTokens.data[token]; exists {
			tokenStoreMu.RUnlock()
			continue
		}
		SingleUseTokens.RUnlock()
		expiresAt := time.Now().Add(2 * time.Minute)
		SingleUseTokens.Lock()
		SingleUseTokens.data[username] = struct {
			Time time.Time; Key string}{expiresAt, token}
		SingleUseTokens.Unlock()
		return token
	}
}

func SingleUseTokenGet(username string) string {
	SingleUseTokens.Lock()
	defer SingleUseTokens.Unlock()
	if tokenData, exists := SingleUseTokens.data[username]; exists {
		if time.Since(tokenData.Time) < 2*time.Minute {
			return tokenData.Key
		}
		delete(SingleUseTokens.data, username)
	}
	return ""
}

func SingleUseTokenInvalidationStart(timer time.Duration) {
	ticker := time.NewTicker(timer)
	defer ticker.Stop()
	for range ticker.C {
		for username, data := range SingleUseTokens.data {
			if time.Now().After(data.Time) {
				SingleUseTokens.Lock()
				delete(SingleUseTokens.data, username)
				SingleUseTokens.Unlock()
			}
		}
	}
}
