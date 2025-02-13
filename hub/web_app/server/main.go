package main


import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4" 
)

var jwtsecret = []byte("123abc@123")

type Claims struct{
	Username string `json:"username"`
	jwt.RegisteredClaims
}

//simulated database
var validUsers = map[string]string{
	"testuser":"password123",//username: password
}

type Credentials struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(w http.ResponseWriter,r *http.Request)
{
	//Only POST Methods are allowed
	if(r.Method != http.MethodPost)
	{
		http.Error(w,"Invalid Request Method",http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials

	//Decode received json data in credentials struct
	err := json.NewDecoder(r.Body).Decode(&creds)

	// Validate the request body
	if err!=nil || creds.Username == "" || creds.Password == ""
	{
		http.Error(w,"Invalid Request Body",http.StatusBadRequest)
	}

	password,ok := validUsers[creds.Username]

	//Ensure the key with the given username exists
	if(!ok)
	{
		http.Error(w,"Invalid username",http.StatusUnauthorized)
		return
	}

	//Ensure that password is correct
	if(password != creds.Password)
	{
		http.Error(w,"Invalid password",http.StatusUnauthorized)
		return
	}

	//Generate JWT Token
	tokenString,err := generateJWT(creds.Username)
	if err!=nil{
		http.Error(w, "Error generating token",http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(map[string]string{"token":tokenString})

}