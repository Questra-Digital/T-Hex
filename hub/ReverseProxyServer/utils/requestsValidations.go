package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"ReverseProxyServer/config"
	"ReverseProxyServer/models"
	"ReverseProxyServer/globalVariables"
)

// Extracts session id from json response
func JSONGetSessId(stream string) (string, error) {
	type BodyMain struct {
		Value struct {
			SessionId string `json:"sessionId"`
		} `json:"value"`
	}
	decoder := json.NewDecoder(strings.NewReader(stream))
	var body BodyMain
	err := decoder.Decode(&body)
	if err != nil {
		return "", err
	}
	return body.Value.SessionId, nil
}

// Gets key and project name. Does logging, and key validation
func GetKeyProjValidate(r *http.Request) (string, string, bool) {
	key := r.Header.Get(globalVariables.HEAD_KEY)
	proj := r.Header.Get(globalVariables.HEAD_NAME)
	if proj == "" {
		proj = globalVariables.HEAD_NAME_DEF
	}
	log.Printf("%s %s; proj: `%s`; key: `%s`", r.Method, r.URL.String(),proj, key)

	if !KeyIsValid(key) {
		log.Printf("\tDropped due to Unauthorized Key: `%s`", key)
		return key, proj, false
	}
	return key, proj, true
}

// Gets key, project name, and test Id. Does logging, and key validation
func GetKeyProjTestIdValidate(r *http.Request) (string, string, int64, bool) {
	key := r.Header.Get(globalVariables.HEAD_KEY)
	proj := r.Header.Get(globalVariables.HEAD_NAME)
	if proj == "" {
		proj = globalVariables.HEAD_NAME_DEF
	}
	testStr := r.Header.Get(globalVariables.HEAD_TEST)
	testId, err := strconv.ParseInt(testStr, 10, 64)
	if err != nil {
		log.Printf("Invalid test Id provided: `%s`", testStr)
		return key, proj, 0, false
	}
	log.Printf("%s %s; proj: `%s`; test: `%d` key: `%s`", r.Method, r.URL.String(),
		proj, testId, key)
	if !KeyIsValid(key) {
		log.Printf("\tDropped due to Unauthorized Key: `%s`", key)
		return key, proj, testId, false
	}
	return key, proj, testId, true
}

// reads body from request
func GetReqBody(r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("\tFailed to parse body: %s", err.Error())
		return nil, err
	}
	r.Body.Close()
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	return body, nil
}

// logs request response
func LogReqRes(r *http.Request, statusCode int,
	res, key, proj, sessId string) error {
	body, err := GetReqBody(r)
	if err != nil {
		return err
	}
	entry := &models.Event{
		Time:      time.Now().Unix(),
		Method:    r.Method,
		Path:      r.URL.Path,
		ReqBody:   string(body),
		Status:    statusCode,
		Res:       string(res),
		SessionId: sessId,
	}
	err = config.DBInit().Create(&entry).Error
	if err != nil {
		log.Fatalf("Failed to log to DB: %s", err.Error())
	}
	log.Printf("\tresponded to: %s %s; proj: `%s`; key: `%s`: %d",
		r.Method, r.URL.String(), proj, key, statusCode)
	return nil
}
