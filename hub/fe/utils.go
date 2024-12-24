package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/mail"
	"sync"
	"time"
	"unicode"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
			time.Now().Before(data.Time) {
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

var keyGenMu sync.Mutex

func ApiKeyGenerate(username string) (string, error) {
	currentKey := UserGetCurrentKey(username)
	if currentKey != "" {
		err := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Delete(&ApiKey{}, "key = ?", currentKey).Error; err != nil {
				return err
			}
			if err := tx.Delete(&UserKey{}, "username = ? AND key = ?", username,
					currentKey).Error; err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return "", fmt.Errorf("error deleting current key: %w", err)
		}
	}

	var newKey string
	keyGenMu.Lock()
	defer keyGenMu.Unlock()
	for {
		keyBytes := make([]byte, 64)
		_, err := rand.Read(keyBytes)
		if err != nil {
			panic("Error generating random key:" + err.Error())
		}
		newKey = hex.EncodeToString(keyBytes)
		var count int64
		err = db.Model(&ApiKey{}).Where("key = ?", newKey).Count(&count).Error
		if err != nil {
			return "", fmt.Errorf("error checking ApiKey table: %w", err)
		}
		if count == 0 {
			err = db.Model(&UserKey{}).Where("key = ?", newKey).Count(&count).Error
			if err != nil {
				return "", fmt.Errorf("error checking UserKey table: %w", err)
			}
		}
		if count == 0 {
			break
		}
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&ApiKey{Key: newKey}).Error; err != nil {
			return err
		}
		if err := tx.Create(&UserKey{Username: username, Key: newKey}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("error inserting new key: %w", err)
	}
	return newKey, nil
}
