package utils

import (
	"cicd_runner/config"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func CreateReverseProxy() (*httputil.ReverseProxy, error) {

	targetURL, err := url.Parse(config.EnvConfig.ReverseProxyUrl)
	if err != nil {
		log.Printf("Error parsing reverse proxy URL: %s", err.Error())
		return nil, err
	}

	//Initially set empty values
	config.ProxyConfig.SetValues("", "", "")

	log.Println("Reverse proxy created - ProxyConfig initialized with empty values")

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Director = func(req *http.Request) {
		// Read ProxyConfig values dynamically for each request
		apiKey, pipelineName, testId := config.ProxyConfig.GetValues()
		req.Header.Add("thex-key", apiKey)
		req.Header.Add("thex-proj", pipelineName)
		req.Header.Add("thex-test", testId)
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		// Strip /reverse_proxy
		req.URL.Path = strings.Replace(req.URL.Path, "/reverse_proxy", "", 1)
		if req.URL.Path == "" {
			req.URL.Path = "/"
		}
	}
	return proxy, nil
}
