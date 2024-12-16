package main

import (
	"html/template"
	"net/http"
	"strconv"
	"embed"

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

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	ClearSessionUser(w, r)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	username := GetSessionUser(r)

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
