package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"time"
)

const ENV_END = "GRID"
const ENV_END_DEF = "http://selenium-hub:4444"

const ENV_LISTEN = "LISTEN"
const ENV_LISTEN_DEF = ":4445"

const ENV_DB = "DB"
const ENV_DB_DEF = "postgres://thex:thex1234@db/thex"

const ENV_TTL = "TTL"
const ENV_TTL_DEF = "120"

const SUB_URL = "/wd/hub"

const HEAD_KEY = "thex-key"
const HEAD_NAME = "thex-proj"
const HEAD_NAME_DEF = "UNNAMED PROJECT"

// TimeToLive, in seconds, for caching keys
var ttl int32

// Get env var, or default val
func Getenv(key string, def string) string {
	ret := os.Getenv(key)
	if ret == "" {
		return def
	}
	return ret
}

func main() {
	endpoint := Getenv(ENV_END, ENV_END_DEF)
	local := Getenv(ENV_LISTEN, ENV_LISTEN_DEF)
	dbStr := Getenv(ENV_DB, ENV_DB_DEF)
	ttlStr := Getenv(ENV_TTL, ENV_TTL_DEF)

	targetURL, err := url.Parse(endpoint)
	if err != nil {
		log.Fatalf("Error parsing THEX_URL `%s`:\n\t%s",
			endpoint, err.Error())
	}

	{
		i, err := strconv.Atoi(ttlStr)
		if err != nil {
			log.Fatalf("Faled to parse TTL: %s", err.Error())
		}
		ttl = int32(i)
	}

	db = DBInit(dbStr)

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.URL.Path = SUB_URL + req.URL.Path
	}

	http.HandleFunc("/", ProxyReqHandler(proxy))
	http.HandleFunc("/session", ProxyReverseHandlerSessionSetup(proxy))
	http.HandleFunc("/session/", ProxyReverseHandlerSession(proxy))

	log.Printf("Starting proxy server:\n\tRemote: `%s`\n\tLocal: `%s`",
		endpoint, local)
	if err := http.ListenAndServe(local, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// Reverse proxy handler for non-/session paths
func ProxyReqHandler(proxy *httputil.ReverseProxy) func(
	http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, valid := GetKeyProjValidate(r)
		if !valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		proxy.ServeHTTP(w, r)
	}
}

// Reverse proxy handler, for setting up session
func ProxyReverseHandlerSessionSetup(proxy *httputil.ReverseProxy) func(
	http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Session Setup Handler")
		key, proj, valid := GetKeyProjValidate(r)
		if !valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		sw := &ResponseWriterSaver{ResponseWriter: w}
		proxy.ServeHTTP(sw, r)
		if r.URL.Path == "/session" && r.Method == "POST" {
			sessId, err := JSONGetSessId(string(sw.body.Bytes()))
			if err != nil {
				log.Printf("\tError in getting New SessionId: %s", err.Error())
			} else {
				db.Create(&KeySession{
					Time:      time.Now().Unix(),
					Key:       key,
					Proj:      proj,
					SessionId: sessId,
					Valid:     true,
				})
				log.Printf("\tNew SessionId: `%s`", sessId)
			}
		} else {
			log.Printf("\tUnsupported method. Only POST supported")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		LogReqRes(r, sw.statusCode, sw.body.String(), key, proj)
	}
}

// reverse proxy handler for all /session/<id>... paths
func ProxyReverseHandlerSession(proxy *httputil.ReverseProxy) func(
	http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Session Handler")
		key, proj, valid := GetKeyProjValidate(r)
		if !valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		sessId, err := URLGetSessId(r.URL.Path)
		if err != nil {
			// TODO: figure out what to actually do here
			log.Printf("\tFailed to extract session ID from URL: %s", err.Error())
			sessId = ""
		}
		// make sure key for this sess is valid
		if !KeySessIsValid(key, sessId) {
			log.Printf("\tDropped due to Key not valid for SessionId")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		sw := &ResponseWriterSaver{ResponseWriter: w}
		proxy.ServeHTTP(sw, r)
		LogReqRes(r, sw.statusCode, sw.body.String(), key, proj)
	}
}
