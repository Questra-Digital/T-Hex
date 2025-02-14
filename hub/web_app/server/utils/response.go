package utils

import (
	"encoding/json"
	"net/http"
)

func RespondJSON(w http.ResponseWriter,data interface{}){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(data)
}

func RespondError(w http.ResponseWriter,message string, statuscode int){
	w.WriteHeader(statuscode)
	RespondJSON(w,map[string]string{"error":message})
}