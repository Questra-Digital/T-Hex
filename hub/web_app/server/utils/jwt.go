package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"server/models"
)

var Jwtsecret = []byte("123abc@123")

func GenerateJWT(username string) (string,error){

	//Expiration is set to 5 days
	expirationTime := time.Now().Add(5 * 24 * time.Hour)

	//Pointer to Claims is created, as jwt library expects a pointer
	claims := &models.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},	
	}

	//Create new token with claims as payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Sign the token with the secret key
	return token.SignedString(Jwtsecret)

}

