package main

import (
	"log"
	"net/http"
	"os"
	"context"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionUsername := GetSessionUser(r)
		if sessionUsername == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r.WithContext(
			context.WithValue(r.Context(), "username", sessionUsername),
		))
	})
}

func main() {
	db = DBInit()
	r := mux.NewRouter()

	r.HandleFunc("/qs.css", ServeCSS).Methods("GET")
	r.HandleFunc("/", ServeLandingPage).Methods("GET")
	r.HandleFunc("/login", LoginHandler).Methods("GET", "POST")
	r.HandleFunc("/signup", SignupHandler).Methods("GET", "POST")
	r.HandleFunc("/contactus", ContactusHandler).Methods("GET", "POST")
	r.HandleFunc("/logout", LogoutHandler).Methods("GET")
	r.Handle("/dashboard", AuthMiddleware(http.HandlerFunc(DashboardHandler)))
	r.Handle("/testsession/{testId:[0-9]+}", AuthMiddleware(http.HandlerFunc(TestSessionHandler)))
	r.Handle("/session/{sessionId}", AuthMiddleware(http.HandlerFunc(SessionHandler)))

	listen := os.Getenv("LISTEN")
	if listen == "" {
		listen = ":8080"
	}

	log.Printf("Server running on port %s", listen)
	log.Fatal(http.ListenAndServe(listen, r))
}
