package main

import (
	"ReverseProxyServer/config"
	"ReverseProxyServer/globalVariables"
	"ReverseProxyServer/models"
	"ReverseProxyServer/utils"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"
)

func main() {
	config.EnvConfigInit()

	endpoint := config.EnvConfig.SeleniumHubUrl
	local := config.EnvConfig.Port

	//Initialize database
	config.DBInit()

	//Parse endpoint URL into targetURL
	targetURL, err := url.Parse(endpoint)
	if err != nil {
		log.Fatalf("Error parsing THEX_URL `%s`:\n\t%s",
			endpoint, err.Error())
	}

	//Setup reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.URL.Path = globalVariables.SUB_URL + req.URL.Path
	}

	//Setup handlers
	http.HandleFunc("/", ProxyReqHandler(proxy))
	http.HandleFunc("/thex/test", THexTestSessSetupHandler)
	http.HandleFunc("/session", ProxyReverseHandlerSessionSetup(proxy))
	http.HandleFunc("/session/", ProxyReverseHandlerSession(proxy))

	//Start server
	if err := http.ListenAndServe(local, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Printf("Starting proxy server:\n\tRemote: %s\n\tLocal: %s", config.EnvConfig.SeleniumHubUrl, config.EnvConfig.Port)
}

// Reverse proxy handler for non-/session paths
func ProxyReqHandler(proxy *httputil.ReverseProxy) func(
	http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, valid := utils.GetKeyProjValidate(r)
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
	key, proj, valid := utils.GetKeyProjValidate(r)
	if !valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if r.Method == "POST" {
		pipelineEventIDStr := r.Header.Get(globalVariables.HEAD_PIPELINE_EVENT_ID)
		pipelineEventID, err := strconv.ParseInt(pipelineEventIDStr, 10, 64)
		if err != nil {
			log.Printf("\tError parsing pipeline event ID '%s': %s", pipelineEventIDStr, err.Error())
			http.Error(w, "Invalid Pipeline Event ID", http.StatusBadRequest)
			return
		}
		testId, err := utils.TestSessCreate(key, proj, pipelineEventID)
		if err != nil {
			log.Printf("\tError setting up test session: %s", err.Error())
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}
		w.Header().Add("thex-test", strconv.FormatInt(testId, 10))
		log.Printf("\tassigned testId: %s %s; proj: `%s`; key: `%s`; testId: `%d",
			r.Method, r.URL.String(), proj, key, testId)
		return
	} else if r.Method == "DELETE" {
		testStr := r.Header.Get(globalVariables.HEAD_TEST)
		testId, err := strconv.ParseInt(testStr, 10, 64)
		if err != nil {
			log.Printf("\tError getting testId: %s", err.Error())
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}

		err = utils.TestSessEnd(testId)
		if err != nil {
			log.Printf("\tError closing test session: %s", err.Error())
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}
		log.Printf("\tclosed test session: %d", testId)
		return
	} else {
		http.Error(w, "bad", http.StatusBadRequest)
	}
}

// Reverse proxy handler, for setting up session
func ProxyReverseHandlerSessionSetup(proxy *httputil.ReverseProxy) func(
	http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Session Setup Handler")
		key, proj, testId, valid := utils.GetKeyProjTestIdValidate(r)
		if !valid || !utils.KeyIsValidForTest(key, testId) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if r.Method != "POST" {
			log.Printf("\tUnsupported method. Only POST supported")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		sw := &utils.ResponseWriterSaver{ResponseWriter: w}
		proxy.ServeHTTP(sw, r)
		sessId, err := utils.JSONGetSessId(sw.Body())
		if err != nil {
			log.Printf("\tError in getting New SessionId: %s", err.Error())
		} else {
			// TODO get current test Id
			config.DB.Create(&models.SessionSelenium{
				Time:      time.Now().Unix(),
				TestId:    testId,
				SessionId: sessId,
				Valid:     true,
			})
			log.Printf("\tNew SessionId: `%s`", sessId)
		}
		utils.LogReqRes(r, sw.StatusCode(), sw.Body(), key, proj, sessId)
	}
}

// reverse proxy handler for all /session/<id>... paths
func ProxyReverseHandlerSession(proxy *httputil.ReverseProxy) func(
	http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Session Handler")
		key, proj, valid := utils.GetKeyProjValidate(r)
		if !valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		sessId, err := utils.URLGetSessId(r.URL.Path)
		if err != nil || !utils.KeyIsValidForSess(key, sessId) {
			log.Printf("\tDropped due to Key not valid for SessionId")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// handle DELETE
		if r.Method == "DELETE" && r.URL.Path == "/session/"+sessId {
			log.Printf("\tDeleting session %s", sessId)

			// Analyze events and set session status before invalidating
			success, err := utils.UpdateSessionStatus(sessId)
			if err != nil {
				log.Printf("\tFailed to update session status: %s", err.Error())
			} else {
				log.Printf("\tSession %s test result: %t", sessId, success)
			}

			// Invalidate the session
			err = utils.KeyMakeInvalidForSess(key, sessId)
			if err != nil {
				log.Printf("\tFailed to Invalidate in DB")
			}
		}

		sw := &utils.ResponseWriterSaver{ResponseWriter: w}
		proxy.ServeHTTP(sw, r)
		utils.LogReqRes(r, sw.StatusCode(), sw.Body(), key, proj, sessId)
	}
}
