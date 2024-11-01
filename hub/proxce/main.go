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
	key := os.Getenv("THEX_KEY")
	if key == "" {
		log.Fatalf("Cannot run with invalid key")
	}
	if endpoint == "" {
		endpoint = "http://localhost:4445"
	}
	if local == "" {
		local = ":4444"
	}
	targetURL, err := url.Parse(endpoint)
	if err != nil {
		log.Printf("Error parsing THEX_URL `%s`:\n\t%s", endpoint, err.Error())
		os.Exit(1)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Director = func(req *http.Request) {
		req.Header.Add("thex-key", key)
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	log.Printf("Starting proxy server:\n\tRemote: `%s`\n\tLocal: `%s`",
		endpoint, local)
	if err := http.ListenAndServe(local, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func KeyIsValid(key string) bool {
	if key == "" {
		return false
	}
	// TODO do the check
	return true
}
