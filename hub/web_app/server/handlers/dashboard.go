package handlers

import (
	"fmt"
	"net/http"
	"server/utils"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request){
	username := r.Header.Get("Username")
	message := fmt.Sprintf("Welcome to the Dashboard, %s!", username)
	utils.RespondJSON(w, http.StatusOK, map[string]string{
		"message": message,
	})
}