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

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const ENV_END = "GRID"
const ENV_END_DEF = "http://selenium-hub:4444"

const ENV_LISTEN = "LISTEN"
const ENV_LISTEN_DEF = ":4445"

const ENV_DB = "DB"
const ENV_DB_DEF = "postgres://thex:thex1234@db/thex"

const ENV_TTL = "TTL"
const ENV_TTL_DEF = "120"

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

	db, err = gorm.Open(postgres.Open(dbStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}
	db.AutoMigrate(&ApiKey{})
	kcache = make(map[string]KCacheEntry)

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.URL.Path = "/wd/hub" + req.URL.Path
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
		key := r.Header.Get("thex-key")
		proj := r.Header.Get("thex-proj")
		if proj == "" {
			proj = "UNNAMED PROJECT"
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

		entry := ReqResPair{
			reqMethod:  r.Method,
			reqPath:    r.URL.Path,
			reqBody:    body,
			reqHeaders: r.Header,
			apiKey:     key,
			proj:       proj,
			resStatus:  sw.statusCode,
			resBody:    sw.body.String(),
		}
		log.Printf("%+v\n\n", entry)
		// TODO: log entry to db
	}
}

// Request and Response pair, for logging
type ReqResPair struct {
	reqMethod  string
	reqPath    string
	reqBody    []byte
	reqHeaders map[string][]string
	apiKey     string
	proj       string
	resStatus  int
	resBody    string
}

type ApiKey struct {
	gorm.Model
	Key string
}

type Tabler interface {
	TableName() string
}

func (ApiKey) TableName() string {
	return "api_keys"
}

// API Key cache entry
type KCacheEntry struct {
	time  int64
	valid bool
}

// TimeToLive, in seconds, for caching keys
var ttl int32

// database
var db *gorm.DB

// api keys cache
var kcache map[string]KCacheEntry

// whether a key is valid or not
// checks the following: key format and cache, and finally DB (if necessary)
func KeyIsValid(key string) bool {
	if key == "" {
		return false
	}
	entry, ok := kcache[key]

	if !ok || time.Now().Unix()-entry.time > int64(ttl) {
		// refresh cache
		exist := false
		_ = db.Model(&ApiKey{}).
			Select("count(*) > 0").
			Where("key = ?", key).
			Find(&exist).
			Error
		kcache[key] = KCacheEntry{time: time.Now().Unix(), valid: exist}
		return exist
	}
	return entry.valid
}
