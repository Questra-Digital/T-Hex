package main

import (
	"embed"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//go:embed templates/*.html
var templatesFS embed.FS

//go:embed qs.css
var cssFS embed.FS

var templates = template.Must(template.ParseFS(templatesFS, "templates/*.html"))

func TemplateExec(w io.Writer, tmpl string, vals map[string]interface{}) {
	if vals == nil {
		templates.ExecuteTemplate(w, tmpl, map[string]interface{}{"Year": time.Now().Year()})
		return
	}
	vals["Year"] = time.Now().Year()
	templates.ExecuteTemplate(w, tmpl, vals)
}

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
	TemplateExec(w, "landing.html", nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		TemplateExec(w, "login.html", nil)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	var user User
	if err := db.First(&user, "username = ?", username).Error; err != nil {
		TemplateExec(w, "login.html", map[string]interface{}{
			"Error": "Invalid credentials"})
		return
	}

	if !PasswordHashMatch(user.Password, password) {
		TemplateExec(w, "login.html", map[string]interface{}{
			"Error": "Invalid credentials"})
		return
	}

	SetSessionUser(w, username)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		TemplateExec(w, "signup.html", nil)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	cpassword := r.FormValue("cpassword")
	if !UsernameIsValid(username) {
		TemplateExec(w, "signup.html", map[string]interface{}{
				"Error": "Username is not valid. " +
					"Username must be alphanumeric, at least 4 characters long.",
			})
		return
	}
	if password != cpassword {
		TemplateExec(w, "signup.html", map[string]interface{}{
				"Error": "Passwords do not match.",
			})
		return
	}
	if UsernameExists(username) {
		TemplateExec(w, "signup.html", map[string]interface{}{
				"Error": "Username is taken.",
			})
		return
	}
	// create
	var user User
	user.Username = username
	var err error
	user.Password, err = PasswordHash(password)
	if err != nil {
		TemplateExec(w, "error.html", nil)
		return
	}
	err = db.Save(&user).Error
	if err != nil {
		TemplateExec(w, "error.html", nil)
		return
	}
	SetSessionUser(w, username)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func ContactusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		TemplateExec(w, "contactus.html", nil)
		return
	}
	email := r.FormValue("email")
	message := r.FormValue("message")
	if !EmailIsValid(email) {
		TemplateExec(w, "contactus.html", map[string]interface{}{
				"Error": "Email is not valid.",
			})
		return
	}
	if message == "" {
		TemplateExec(w, "contactus.html", map[string]interface{}{
				"Error": "Please fill the Message field",
			})
		return
	}
	var msg ContactUsMessage
	msg.Email = email
	msg.Message = message
	if err := db.Save(&msg).Error; err != nil {
		TemplateExec(w, "error.html", nil)
		return
	}
	TemplateExec(w, "contactus.html", map[string]interface{}{
		"Message": "Thank you for your message",
	})
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

	TemplateExec(w, "dashboard.html", map[string]interface{}{
		"TestSessions": testSessions,
	})
}

func ApiKeyHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	token := SingleUseTokenGenerate(username)
	if token == "" {
		TemplateExec(w, "apikeys.html", map[string]interface{}{
			"Error": "You have already visited this page within last 2 minutes. " +
			"Try again later",
		})
	}
	TemplateExec(w, "apikeys.html", map[string]interface{}{
		"Token": token,
	})
}

func ApiKeyGenHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	token := r.PostFormValue("token")
	if username == "" || token == "" || SingleUseTokenGet(username) != token {
		http.Error(w, "Bad Token", http.StatusUnauthorized)
		return
	}
	key, err := ApiKeyGenerate(username)
	if err != nil {
		TemplateExec(w, "apikeyres.html", map[string]interface{}{
			"key": "ERROR GENERATING KEY",
		})
		return
	}
	TemplateExec(w, "apikeyres.html", map[string]interface{}{
		"key": key,
	})
}

func TestSessionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	testId, _ := strconv.ParseInt(vars["testId"], 10, 64)

	var sessions []Session
	db.Where("test_id = ?", testId).Order("time DESC").Find(&sessions)

	TemplateExec(w, "testsession.html", map[string]interface{}{
		"Sessions": sessions,
	})
}

func SessionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionId := vars["sessionId"]

	var events []Event
	db.Where("session_id = ?", sessionId).Order("time DESC").Find(&events)

	TemplateExec(w, "session.html", map[string]interface{}{
		"Events": events,
	})
}
