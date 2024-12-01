package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
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

type StringInt struct {
	A string
	B int64
}

// cache for Key to Session Id mapping
var ksessCache map[StringPair]KSessCacheEntry =
	make(map[StringPair]KSessCacheEntry)

// whether a key is valid for a session or not
func KeyIsValidForSess(key string, sessId string) bool {
	if key == "" || sessId == "" {
		return false
	}
	entry, ok := ksessCache[StringPair{key, sessId}]

	if !ok || time.Now().Unix()-entry.Time > int64(ttl) {
		// refresh cache
		var result bool
		err := db.Table("sessions").
			Select("COUNT(*) > 0").
			Joins("JOIN test_sessions ON sessions.test_id = test_sessions.test_id").
			Where("sessions.session_id = ?", sessId).
			Where("sessions.valid = ?", true).
			Where("test_sessions.key = ?", key).
			Where("test_sessions.current = ?", true).
			Find(&result).Error
		exist := err == nil && result
		ksessCache[StringPair{key, sessId}] =
			KSessCacheEntry{Time: time.Now().Unix(), Valid: exist}
		return exist
	}
	return entry.Valid
}

// cache for Key to test Id mapping
var ktestCache map[StringInt]KSessCacheEntry =
	make(map[StringInt]KSessCacheEntry)

// whether a key is valid for a test or not
func KeyIsValidForTest(key string, testId int64) bool {
	if key == "" {
		return false
	}
	entry, ok := ktestCache[StringInt{key, testId}]

	if !ok || time.Now().Unix()-entry.Time > int64(ttl) {
		// refresh cache
		var result bool
		err := db.Table("test_sessions").
			Select("COUNT(*) > 0").
			Where("test_id = ?", testId).
			Where("key = ?", key).
			Where("current = ?", true).
			Find(&result).Error
		exist := err == nil && result
		ktestCache[StringInt{key, testId}] =
			KSessCacheEntry{Time: time.Now().Unix(), Valid: exist}
		return exist
	}
	return entry.Valid
}

// makes a key valid for a session id, with a project name
func KeyMakeValidForSess(testId int64, key string, sessId string, proj string,
		) error {
	if key == "" || sessId == "" {
		return errors.New("key or session Id is empty")
	}
	cacheEntry := KSessCacheEntry{
		Time: time.Now().Unix(),
		Valid: true,
	}
	ksessCache[StringPair{key, sessId}] = cacheEntry
	entry := Session{
		Time: cacheEntry.Time,
		TestId: testId,
		SessionId: sessId,
		Valid: true,
	}
	err := db.Create(entry).Error
	if err != nil {
		return err
	}
	return nil
}

// makes a key invalid for a session id
func KeyMakeInvalidForSess(key string, sessId string) error {
	if key == "" || sessId == "" {
		return errors.New("key or session Id is empty")
	}
	// invalidate in cache
	ksessCache[StringPair{key, sessId}] = KSessCacheEntry{0, false}
	// invalidate in db
	err := db.Model(&Session{}).
		Where("session_id = ?", sessId).
		Update("valid", false).
		Error
	if err != nil {
		return err
	}
	return nil
}

// makes a new Testing Session for a key and proj
func TestSessCreate(key string, proj string) (int64, error) {
	test := TestSession{
		Time: time.Now().Unix(),
		Key: key,
		Proj: proj,
		Current: true,
	}
	err := db.Create(&test).Error
	if err != nil {
		return 0, err
	}
	return test.TestId, nil
}

// ends a testing session
func TestSessEnd(testId int64) error {
	err := db.Model(&TestSession{}).
		Where("test_id = ?", testId).
		Update("current", false).
		Error
	if err != nil {
		return err
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

// Extracts test id from url
func URLGetTestId(URL string) (int64, error) {
	if !strings.HasPrefix(URL, "/thex/test/") {
		return 0, errors.New("Not a /thex/test/<id> url")
	}
	s, _ := strings.CutPrefix(URL, "/session/")
	s, _, _ = strings.Cut(s, "/")
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
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

// Gets key and project name. Does logging, and key validation
func GetKeyProjValidate(r *http.Request) (string, string, bool) {
	key := r.Header.Get(HEAD_KEY)
	proj := r.Header.Get(HEAD_NAME)
	if proj == "" {
		proj = HEAD_NAME_DEF
	}
	log.Printf("%s %s; proj: `%s`; key: `%s`", r.Method, r.URL.String(),
		proj, key)
	if !KeyIsValid(key) {
		log.Printf("\tDropped due to Unauthorized Key: `%s`", key)
		return key, proj, false
	}
	return key, proj, true
}

// Gets key, project name, and test Id. Does logging, and key validation
func GetKeyProjTestIdValidate(r *http.Request) (string, string, int64, bool) {
	key := r.Header.Get(HEAD_KEY)
	proj := r.Header.Get(HEAD_NAME)
	if proj == "" {
		proj = HEAD_NAME_DEF
	}
	testStr := r.Header.Get(HEAD_TEST)
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
	entry := &Event{
		Time:    time.Now().Unix(),
		Method:  r.Method,
		Path:    r.URL.Path,
		ReqBody: string(body),
		Status:  statusCode,
		Res:     string(res),
		SessionId: sessId,
	}
	err = db.Create(&entry).Error
	if err != nil {
		log.Fatalf("Failed to log to DB: %s", err.Error())
	}
	log.Printf("\tresponded to: %s %s; proj: `%s`; key: `%s`: %d",
		r.Method, r.URL.String(), proj, key, statusCode)
	return nil
}
