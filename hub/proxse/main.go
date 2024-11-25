package main

import (
	"bytes"
	"io"
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

// api keys cache
var kcache map[string]KCacheEntry

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
	kcache = make(map[string]KCacheEntry)

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.URL.Path = SUB_URL + req.URL.Path
	}

	http.HandleFunc("/", proxyReqHandler(proxy))

	log.Printf("Starting proxy server:\n\tRemote: `%s`\n\tLocal: `%s`",
		endpoint, local)
	if err := http.ListenAndServe(local, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// Reverse proxy handler
func proxyReqHandler(proxy *httputil.ReverseProxy) func(
		http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get(HEAD_KEY)
		proj := r.Header.Get(HEAD_NAME)
		if proj == "" {
			proj = HEAD_NAME_DEF
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf(
				"Erroneous Request: %s %s Headers:\n\t%v\n\n",
				r.Method, r.URL.Path, r.Header)
			log.Printf("Error reading body")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		r.Body.Close()
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		if !KeyIsValid(key) {
			log.Print("Dropped due to Unauthorized")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		sw := &ResponseWriterSaver{ResponseWriter: w}
		proxy.ServeHTTP(sw, r)

		entry := &EventLogEntry{
			time: time.Now().Unix(),
			method: r.Method,
			path: r.URL.Path,
			reqBody: string(body),
			key: key,
			proj: proj,
			status: sw.statusCode,
			res: sw.body.String(),
		}
		log.Printf("%+v\n\n", entry)
		err = EventLogToDB(entry)
		if err != nil {
			log.Fatalf("Failed to log to DB: %s", err.Error())
		}
	}
}
