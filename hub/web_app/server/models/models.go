package models

import "github.com/golang-jwt/jwt/v4"

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