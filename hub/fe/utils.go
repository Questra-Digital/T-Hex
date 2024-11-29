package main

import (
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
