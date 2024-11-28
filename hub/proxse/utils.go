package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// API Key cache entry
type KCacheEntry struct {
	time  int64
	valid bool
}

// api keys cache
var keycache map[string]KCacheEntry = make(map[string]KCacheEntry)

// whether a key is valid or not
func KeyIsValid(key string) bool {
	if key == "" {
		return false
	}
	entry, ok := keycache[key]

	if !ok || time.Now().Unix()-entry.time > int64(ttl) {
		// refresh cache
		exist := false
		_ = db.Model(&ApiKey{}).
			Select("count(*) > 0").
			Where("key = ?", key).
			Find(&exist).
			Error
		keycache[key] = KCacheEntry{time: time.Now().Unix(), valid: exist}
		return exist
	}
	return entry.valid
}

// API Key to Session Id cache entry
type KSessCacheEntry struct {
	Time  int64
	Valid bool
}

type StringPair struct {
	A string
	B string
}

// cache for Key to Session Id mapping
var ksessCache map[StringPair]KSessCacheEntry =
	make(map[StringPair]KSessCacheEntry)

// whether a key is valid for a session or not
func KeySessIsValid(key string, sessId string) bool {
	if key == "" || sessId == "" {
		return false
	}
	entry, ok := ksessCache[StringPair{key, sessId}]

	if !ok || time.Now().Unix()-entry.Time > int64(ttl) {
		// refresh cache
		exist := false
		_ = db.Model(&KeySession{}).
			Select("count(*) > 0").
			Where("key = ?", key).
			Where("session_id = ?", sessId).
			Where("valid = ?", true).
			Find(&exist).
			Error
		ksessCache[StringPair{key, sessId}] =
			KSessCacheEntry{Time: time.Now().Unix(), Valid: exist}
		return exist
	}
	return entry.Valid
}

// makes a key valid for a session id, with a project name
func KeySessMakeValid(key string, sessId string, proj string) error {
	if key == "" || sessId == "" {
		return errors.New("key or session Id is empty")
	}
	cacheEntry := KSessCacheEntry{
		Time: time.Now().Unix(),
		Valid: true,
	}
	ksessCache[StringPair{key, sessId}] = cacheEntry
	entry := KeySession{
		Time: cacheEntry.Time,
		Key: key,
		Proj: proj,
		SessionId: sessId,
		Valid: true,
	}
	res := db.Create(entry)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// makes a key invalid for a session id
func KeySessMakeInvalid(key string, sessId string) error {
	if key == "" || sessId == "" {
		return errors.New("key or session Id is empty")
	}
	// invalidate in cache
	ksessCache[StringPair{key, sessId}] = KSessCacheEntry{0, false}
	// invalidate in db
	err := db.Model(&KeySession{}).
		Where("key = ?", key).
		Where("session_id = ?", sessId).
		Update("valid", false).
		Error
	if err != nil {
		return err
	}
	return nil
}

// Inserts a log entry
func EventLogToDB(event *EventLogEntry) error {
	res := db.Create(event)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// Extracts session id from url
func URLGetSessId(URL string) (string, error) {
	if !strings.HasPrefix(URL, "/session/") {
		return "", errors.New("Not a /session/<id> url")
	}
	s, _ := strings.CutPrefix(URL, "/session/")
	s, _, _ = strings.Cut(s, "/")
	return s, nil
}

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

// Returns: (key, project name)
func ReqGetKeyProj(r *http.Request) (string, string) {
	key := r.Header.Get(HEAD_KEY)
	proj := r.Header.Get(HEAD_NAME)
	if proj == "" {
		proj = HEAD_NAME_DEF
	}
	return key, proj
}

// Gets key and project name. Does logging, and key validation
func GetKeyProjValidate(r *http.Request) (string, string, bool) {
	key, proj := ReqGetKeyProj(r)
	log.Printf("%s %s; proj: `%s`; key: `%s`", r.Method, r.URL.String(),
	proj, key)
	if !KeyIsValid(key) {
		log.Printf("\tDropped due to Unauthorized Key: `%s`", key)
		return key, proj, false
	}
	return key, proj, true
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
func LogReqRes(r *http.Request, statusCode int, res string, key string,
		proj string) error {
	body, err := GetReqBody(r)
	if err != nil {
		return err
	}
	entry := &EventLogEntry{
		Time:    time.Now().Unix(),
		Method:  r.Method,
		Path:    r.URL.Path,
		ReqBody: string(body),
		Key:     key,
		Proj:    proj,
		Status:  statusCode,
		Res:     string(res),
	}
	err = EventLogToDB(entry)
	if err != nil {
		log.Fatalf("Failed to log to DB: %s", err.Error())
	}
	log.Printf("\tresponded to: %s %s; proj: `%s`; key: `%s`: %d",
		r.Method, r.URL.String(), proj, key, statusCode)
	return nil
}
