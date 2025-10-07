package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const ENV_END = "GRID"
const ENV_END_DEF = "http://selenium-hub:4444"

const ENV_LISTEN = "LISTEN"
const ENV_LISTEN_DEF = ":4445"

const ENV_TTL = "TTL"
const ENV_TTL_DEF = "120"

const ENV_NODB = "THEX_NODB"
const ENV_NODB_DEF = "DB"

const SUB_URL = "/wd/hub"

const HEAD_KEY = "thex-key"
const HEAD_NAME = "thex-proj"
const HEAD_TEST = "thex-test"
const HEAD_NAME_DEF = "UNNAMED PROJECT"

// TimeToLive, in seconds, for caching keys
var ttl int32

var db *gorm.DB

var noDb bool; // if running under no database mode.

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
	ttlStr := Getenv(ENV_TTL, ENV_TTL_DEF)
	noDb = Getenv(ENV_NODB, ENV_NODB_DEF) != ENV_NODB_DEF;

	if noDb {
		if isDummyGrid(endpoint) {
			log.Println("Failed to detect dummy grid. Will not run in NODB mode!");
			noDb = false
		} else {
			log.Println("RUNNING UNDER THEX_NODB. DO NOT USE THIS OUTSIDE TESTING.")
			log.Println("FOLLOWING IS VALID AUTHENTICATION:")
			log.Println("\tUsername:\tnodb")
			log.Println("\tAPI KEY:\tabcd1234")
		}
	}

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

	if noDb{
		db = DBInitNoDB()
	} else {
		db = DBInit()
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.URL.Path = SUB_URL + req.URL.Path
	}

	http.HandleFunc("/", ProxyReqHandler(proxy))
	http.HandleFunc("/thex/test", THexTestSessSetupHandler)
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

// thex/test, for setting up test session
func THexTestSessSetupHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/thex/test Handler")
	key, proj, valid := GetKeyProjValidate(r)
	if !valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if r.Method == "POST" {
		testId, err := TestSessCreate(key, proj)
		if err != nil {
			log.Printf("\tError setting up test session: %s", err.Error())
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}
		w.Header().Add("thex-test", strconv.FormatInt(testId, 10))
		log.Printf("\tassigned testId: %s %s; proj: `%s`; key: `%s`; testId: `%d",
			r.Method, r.URL.String(), proj, key, testId)
		return
	} else
	if r.Method == "DELETE" {
		// TODO implement closing session
	} else {
		http.Error(w, "bad", http.StatusBadRequest)
	}
}

// Reverse proxy handler, for setting up session
func ProxyReverseHandlerSessionSetup(proxy *httputil.ReverseProxy) func(
	http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Session Setup Handler")
		key, proj, testId, valid := GetKeyProjTestIdValidate(r)
		if !valid || !KeyIsValidForTest(key, testId){
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if r.Method != "POST" {
			log.Printf("\tUnsupported method. Only POST supported")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		sw := &ResponseWriterSaver{ResponseWriter: w}
		proxy.ServeHTTP(sw, r)
		sessId, err := JSONGetSessId(string(sw.body.Bytes()))
		if err != nil {
			log.Printf("\tError in getting New SessionId: %s", err.Error())
		} else {
			// TODO get current test Id
			db.Create(&SessionSelenium{
				Time:      time.Now().Unix(),
				TestId:    testId,
				SessionId: sessId,
				Valid:     true,
			})
			log.Printf("\tNew SessionId: `%s`", sessId)
		}
		LogReqRes(r, sw.statusCode, sw.body.String(), key, proj, sessId)
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
		if err != nil || !KeyIsValidForSess(key, sessId) {
			log.Printf("\tDropped due to Key not valid for SessionId")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// handle DELETE
		if r.Method == "DELETE" && r.URL.Path == "/session/" + sessId {
			log.Printf("\tDeleting session %s", sessId)
			err := KeyMakeInvalidForSess(key, sessId)
			if err != nil {
				log.Printf("\tFailed to Invalidate in DB")
			}
		}

		sw := &ResponseWriterSaver{ResponseWriter: w}
		proxy.ServeHTTP(sw, r)
		LogReqRes(r, sw.statusCode, sw.body.String(), key, proj, sessId)
	}
}

// initializes DB for THEX_NODB
func DBInitNoDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Println("Failed to initialize dummy DB:", err)
		return db
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	err = db.AutoMigrate(&UserKey{})
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	err = db.AutoMigrate(&Event{})
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	err = db.AutoMigrate(&SessionSelenium{})
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	err = db.AutoMigrate(&TestSession{})
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	err = db.AutoMigrate(&ContactUsMessage{})
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	if err = db.Create(&User{"nodb", "DUMMY STRING"}).Error; err != nil {
		log.Fatalf("Error: %s", err.Error());
	}
	if err = db.Create(&UserKey{"nodb", "abcd1234"}).Error; err != nil {
		log.Fatalf("Error: %s", err.Error());
	}

	log.Println("Dummy Gorm DB initialized")
	return db
}

func isDummyGrid(endpoint string) bool {
	resp, err := http.Get(endpoint + "/is-dummy-grid")
	if err != nil {
		return true
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return true
	}
	return string(body) == "true\n"
}
