package utils

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"

	"cicd_runner/config"
)

// StartReverseProxy starts a reverse proxy server that stops after 30s of inactivity.
// This call will block until the server stops.
func StartReverseProxy(APIKey, PipelineName, TestId string) {
	targetURL, err := url.Parse(config.EnvConfig.ReverseProxyUrl)
	if err != nil {
		log.Fatalf("Error parsing reverse proxy URL: %s", err.Error())
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Director = func(req *http.Request) {
		req.Header.Set("thex-key", APIKey)
		req.Header.Set("thex-proj", PipelineName)
		req.Header.Set("thex-test", TestId)
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
	}

	server := &http.Server{
		Addr: ":8888",
	}

	// Track last request time
	var mu sync.Mutex
	lastRequest := time.Now()

	// Idle timeout checker
	stopCh := make(chan struct{})
	go func() {
		for {
			time.Sleep(time.Second)
			mu.Lock()
			if time.Since(lastRequest) > 30*time.Second {
				mu.Unlock()
				log.Println("No activity for 30s. Shutting down proxy.")
				server.Shutdown(context.Background())
				close(stopCh)
				return
			}
			mu.Unlock()
		}
	}()

	// Handle requests
	http.HandleFunc("/*path", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		lastRequest = time.Now()
		mu.Unlock()
		proxy.ServeHTTP(w, r)
	})

	// Start server (blocking until Shutdown is called)
	log.Println("Starting reverse proxy on :8888")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Proxy server failed: %v", err)
	}

	// Wait for stop signal
	<-stopCh
	log.Println("Proxy stopped.")
}
