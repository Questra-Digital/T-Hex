package main

import (
	"net/mail"
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
