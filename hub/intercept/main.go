package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// selenium grid URL
	targetURL, _ := url.Parse("http://localhost:4444")

	// Set up the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Director = func(req *http.Request) {
		log.Printf("Intercepted request: %s %s", req.Method, req.URL.Path)
		log.Printf("Headers: %v", req.Header)

		// TODO: remove this
		if req.Method == "DELETE" {
			return
		}

		// forward
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	log.Println("Starting proxy server on :8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
