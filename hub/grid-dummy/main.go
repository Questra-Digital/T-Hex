package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

func randomID() string {
	rand.Seed(0)
	letters := "abcdef0123456789"
	id := make([]byte, 32)
	for i := range id {
		id[i] = letters[rand.Intn(len(letters))]
	}
	return string(id)
}

func main() {
	local := os.Getenv("LOCAL")
	if local == "" {
		local = ":4444"
		log.Printf("Using default `%s` for listener", local)
		log.Println("\tConsider setting the LOCAL env var")
	}
	http.HandleFunc("/session", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			sessionID := randomID()
			response := map[string]any{
				"value": map[string]any{
					"sessionId": sessionID,
					"capabilities": map[string]any{
						"browserName": "chrome",
						"browserVersion": "132.0.6834.159",
						"platformName": "linux",
					},
				},
			}
			json.NewEncoder(w).Encode(response)
		}
	})

	http.HandleFunc("/session/", func(w http.ResponseWriter, r *http.Request) {
		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) < 3 {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		//sessionID := pathParts[2]

		if r.Method == http.MethodDelete {
			json.NewEncoder(w).Encode(map[string]any{"value": nil})
			return
		}

		if strings.HasSuffix(r.URL.Path, "/url") || strings.HasSuffix(r.URL.Path, "/execute/sync") {
			json.NewEncoder(w).Encode(map[string]any{"value": true})
			return
		}

		if strings.HasSuffix(r.URL.Path, "/title") {
			json.NewEncoder(w).Encode(map[string]any{"value": "Example Domain"})
			return
		}

		if strings.HasSuffix(r.URL.Path, "/element") {
			elementID := randomID()
			json.NewEncoder(w).Encode(map[string]any{
				"value": map[string]string{
					"element-6066-11e4-a52e-4f735466cecf": elementID,
				},
			})
			return
		}

		if strings.Contains(r.URL.Path, "/text") {
			json.NewEncoder(w).Encode(map[string]any{"value": "More information..."})
			return
		}

		json.NewEncoder(w).Encode(map[string]any{"value": nil})
	})

	log.Printf("Selenium dummy hub running on %s\n", local)
	http.ListenAndServe(local, nil)
}
