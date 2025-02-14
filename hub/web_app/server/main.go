package main


import (
	"fmt"
	"log"
	"server/handlers"
	"server/middleware"
	"net/http" 
)

func main(){
	http.HandleFunc("/login",handlers.LoginHandler)
	http.HandleFunc("/dashboard",middleware.AuthMiddleware(handlers.DashboardHandler))
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}