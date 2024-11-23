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

func main() {
	endpoint := os.Getenv("THEX_URL")
	local := os.Getenv("LOCAL")
	dbStr := os.Getenv("DB_CONSTR")
	ttlStr := os.Getenv("TTL")
	if dbStr == "" {
		dbStr = "postgres://thex:thex1234@db/thex"
	}
	if endpoint == "" {
		endpoint = "http://selenium-hub:4444"
	}
	if local == "" {
		local = ":4445"
	}
	if ttlStr == "" {
		ttlStr = "120" // 2 minutes
	}

	targetURL, err := url.Parse(endpoint)
	if err != nil {
		log.Fatalf("Error parsing THEX_URL `%s`:\n\t%s", endpoint, err.Error())
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	})

	log.Printf("Starting proxy server:\n\tRemote: `%s`\n\tLocal: `%s`",
		endpoint, local)
	if err := http.ListenAndServe(local, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
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
