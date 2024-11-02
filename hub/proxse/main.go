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
	dbStr := os.Getenv("DB_CONSTR")
	if dbStr == "" {
		dbStr = "postgres://thex:thex1234@db/thex"
	}
	if endpoint == "" {
		endpoint = "http://selenium-hub:4444"
	}
	if local == "" {
		local = ":4445"
	}
	targetURL, err := url.Parse(endpoint)
	if err != nil {
		log.Fatalf("Error parsing THEX_URL `%s`:\n\t%s", endpoint, err.Error())
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.URL.Path = "/wd/hub" + req.URL.Path
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: log to DB
		key := r.Header.Get("thex-key")
		proj := r.Header.Get("thex-proj")
		if proj == "" {
			proj = "UNNAMED PROJECT"
		}
		log.Printf(
			"Request received: %s %s\n\tkey: `%s`\n\tproj: `%s`\n\tHeaders: %v\n",
			r.Method, r.URL.Path, key, proj, r.Header,
		)

		if !KeyIsValid(key) {
			log.Print("Dropped due to Unauthorized")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

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
	// TODO check in DB
	return true
}
