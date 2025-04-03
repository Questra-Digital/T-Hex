package models

import "github.com/golang-jwt/jwt/v4"
import "sync"

type Claims struct{
	Username string `json:"username"`
	jwt.RegisteredClaims
}

//simulated database
var ValidUsers = map[string]string{
	"testuser":"password123",//username: password
}

type Credentials struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

// In-memory state storage
var StateStore = struct {
	sync.RWMutex
	Map map[string]string
}{Map: make(map[string]string)}