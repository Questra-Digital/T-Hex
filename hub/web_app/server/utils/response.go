package utils

import (
	"encoding/json"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode) // Set the HTTP status code
	json.NewEncoder(w).Encode(data)
}


func RespondError(w http.ResponseWriter,message string, statuscode int){
	RespondJSON(w,statuscode,map[string]string{"error":message})
}