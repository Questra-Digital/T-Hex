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

func loginHandler(w http.ResponseWriter,r *http.Request){
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

func generateJWT(username string) (string,error){

	//Expiration is set to 5 days
	expirationTime := time.Now().Add(5 * 24 * time.Hour)

	//Pointer to Claims is created, as jwt library expects a pointer
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},	
	}

	//Create new token with claims as payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Sign the token with the secret key
	return token.SignedString(jwtsecret)

}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		//Extract token from Authorization Header Field
		tokenString := r.Header.Get("Authorization")

		if tokenString == ""{
			http.Error(w,"Missing Token",http.StatusUnauthorized)
			return
		}

		//Parse and validate the token
		claims := &Claims{}
		token,err := jwt.ParseWithClaims(tokenString,claims,func (token *jwt.Token) (interface{},error){
			return jwtsecret, nil
		})

		//Set request header for subsequent handlers
		r.Header.Set("Username",claims.Username)

		//Call the next handler
		next(w,r)

	}
}

func dashboardHandler(w http.ResponseWriter, r *http.Request)
{
	username := r.Header.Get("Username")
	fmt.Fprintf(w,"Welcome to the Dashboard, %s!",username)
}

func main(){
	http.HandleFunc("/login",loginHandler)
	http.HandleFunc("/dashboard",authMiddleware(dashboardHandler))
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}