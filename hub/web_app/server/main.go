package main

import (
	"fmt"
	"log"
	"server/handlers"
	"server/middleware"
	"net/http" 
	"server/config"
)

func main(){

	config.LoadEnv()

	http.HandleFunc("/login",handlers.LoginHandler)
	http.HandleFunc("/dashboard",middleware.AuthMiddleware(handlers.DashboardHandler))
	http.HandleFunc("/googleLogin",handlers.GoogleLogin)
	http.HandleFunc("/auth/google/callback",handlers.GoogleCallback)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}