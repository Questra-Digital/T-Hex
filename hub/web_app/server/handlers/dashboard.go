package handlers

import (
	"fmt"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request){
	username := r.Header.Get("Username")
	fmt.Fprintf(w,"Welcome to the Dashboard, %s!",username)
}