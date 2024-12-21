package main

import (
	"embed"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//go:embed templates/*.html
var templatesFS embed.FS

//go:embed qs.css
var cssFS embed.FS

var templates = template.Must(template.ParseFS(templatesFS, "templates/*.html"))

func ServeCSS(w http.ResponseWriter, r *http.Request) {
	cssFile, err := cssFS.ReadFile("qs.css")
	if err != nil {
		http.Error(w, "Error loading CSS file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/css")
	w.Write(cssFile)
}

func ServeLandingPage(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "landing.html",
		"© 2024 T-Hex. All rights reserved.")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templates.ExecuteTemplate(w, "login.html", nil)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	var user User
	if err := db.First(&user, "username = ?", username).Error; err != nil {
		templates.ExecuteTemplate(w, "login.html", "Invalid credentials")
		return
	}

	if !PasswordHashMatch(user.Password, password) {
		templates.ExecuteTemplate(w, "login.html", "Invalid credentials")
		return
	}

	SetSessionUser(w, username)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templates.ExecuteTemplate(w, "signup.html", nil)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	cpassword := r.FormValue("cpassword")
	if !UsernameIsValid(username) {
		templates.ExecuteTemplate(w, "signup.html", "Username is not valid. " +
			"Username must be alphanumeric, at least 4 characters long.")
		return
	}
	if password != cpassword {
		templates.ExecuteTemplate(w, "signup.html", "Passwords do not match.")
		return
	}
	if UsernameExists(username) {
		templates.ExecuteTemplate(w, "signup.html", "Username is taken.")
		return
	}
	// create
	var user User
	user.Username = username
	var err error
	user.Password, err = PasswordHash(password)
	if err != nil {
		templates.ExecuteTemplate(w, "error.html", nil)
		return
	}
	err = db.Save(&user).Error
	if err != nil {
		templates.ExecuteTemplate(w, "error.html", nil)
		return
	}
	SetSessionUser(w, username)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func ContactusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templates.ExecuteTemplate(w, "contactus.html", nil)
		return
	}
	email := r.FormValue("email")
	message := r.FormValue("message")
	if !EmailIsValid(email) {
		templates.ExecuteTemplate(w, "contactus.html", "Email is not valid")
		return
	}
	if message == "" {
		templates.ExecuteTemplate(w, "contactus.html", "Please fill the Message field")
		return
	}
	var msg ContactUsMessage
	msg.Email = email
	msg.Message = message
	if err := db.Save(&msg).Error; err != nil {
		templates.ExecuteTemplate(w, "error.html", nil)
		return
	}
	templates.ExecuteTemplate(w, "contactus.html", "Thank you for your message")
	return
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	ClearSessionUser(w, r)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username")

	var userKeys []UserKey
	db.Where("username = ?", username).Find(&userKeys)

	var testSessions []TestSession
	for _, key := range userKeys {
		var sessions []TestSession
		db.Where("key = ?", key.Key).Order("time DESC").Find(&sessions)
		testSessions = append(testSessions, sessions...)
	}

	templates.ExecuteTemplate(w, "dashboard.html", testSessions)
}

func TestSessionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	testId, _ := strconv.ParseInt(vars["testId"], 10, 64)

	var sessions []Session
	db.Where("test_id = ?", testId).Order("time DESC").Find(&sessions)

	templates.ExecuteTemplate(w, "testsession.html", sessions)
}

func SessionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionId := vars["sessionId"]

	var events []Event
	db.Where("session_id = ?", sessionId).Order("time DESC").Find(&events)

	templates.ExecuteTemplate(w, "session.html", events)
}
