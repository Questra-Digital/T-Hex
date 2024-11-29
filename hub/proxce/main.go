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
	proj := os.Getenv("THEX_PROJ")
	if key == "" {
		log.Fatalf("THEX_KEY env var is not set.")
	}
	if endpoint == "" {
		endpoint = "http://localhost:4445"
		log.Printf("Using default `%s` endpoint", endpoint)
		log.Println("\tConsider setting the THEX_URL env var")
	}
	if local == "" {
		local = ":4444"
		log.Printf("Using default `%s` for listener", local)
		log.Println("\tConsider setting the LOCAL env var")
	}
	if proj == "" {
		proj = "UNNAMED PROJECT"
		log.Printf("Using default `%s` for name.", proj)
		log.Println("\tConsider setting the THEX_PROJ env var")
	}
	targetURL, err := url.Parse(endpoint)
	if err != nil {
		log.Fatalf("Error parsing THEX_URL `%s`:\n\t%s", endpoint, err.Error())
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Director = func(req *http.Request) {
		req.Header.Add("thex-key", key)
		req.Header.Add("thex-proj", proj)
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.String())
		proxy.ServeHTTP(w, r)
	})

	log.Printf("Starting proxy server:\n\tRemote: `%s`\n\tLocal: `%s`\n\tProject: `%s`",
		endpoint, local, proj)
	if err := http.ListenAndServe(local, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
