package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	endpoint := os.Getenv("THEX_URL")
	local := os.Getenv("LOCAL")
	if endpoint == "" {
		endpoint = "http://localhost:4444"
	}
	if local == "" {
		local = ":4445"
	}
	targetURL, err := url.Parse(endpoint)
	if err != nil {
		log.Printf("Error parsing THEX_URL `%s`:\n\t%s", endpoint, err.Error())
		os.Exit(1)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Director = func(req *http.Request) {

		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: log to DB
		key := r.Header.Get("thex-key")
		proj := r.Header.Get("thex-proj")
		log.Printf("key: `%s`, proj: `%s`", key, proj)

		if !KeyIsValid(key) {
			log.Print("Dropped due to Unauthorized")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		proxy.ServeHTTP(w, r)
	})

	log.Println("Starting proxy server on :8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func KeyIsValid(key string) bool {
	if key == "" {
		return false
	}
	// TODO check in DB
	return true
}
