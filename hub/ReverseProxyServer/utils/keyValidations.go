package utils

import (
	"ReverseProxyServer/config"
	"ReverseProxyServer/models"
	"errors"
	"time"
)

// API Key cache entry
type KCacheEntry struct {
	time  int64
	valid bool
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
var ksessCache map[StringPair]KSessCacheEntry = make(map[StringPair]KSessCacheEntry)

// api keys cache
var keycache map[string]KCacheEntry = make(map[string]KCacheEntry)

// whether a key is valid or not
func KeyIsValid(key string) bool {
	if key == "" {
		return false
	}
	entry, ok := keycache[key]

	if !ok || time.Now().Unix()-entry.time > int64(config.EnvConfig.CacheTTL) {
		// refresh cache
		exist := false
		_ = config.DB.Model(&models.Keys{}).
			Select("count(*) > 0").
			Where("api_key = ?", key).
			Find(&exist).
			Error
		keycache[key] = KCacheEntry{time: time.Now().Unix(), valid: exist}
		return exist
	}
	return entry.valid
}

// whether a key is valid for a session or not
func KeyIsValidForSess(key string, sessId string) bool {
	if key == "" || sessId == "" {
		return false
	}
	entry, ok := ksessCache[StringPair{key, sessId}]

	if !ok || time.Now().Unix()-entry.Time > int64(config.EnvConfig.CacheTTL) {
		// refresh cache
		var result bool
		err := config.DB.Model(&models.SessionSelenium{}).
			Select("COUNT(*) > 0").
			Joins("JOIN test_sessions ON sessions_selenium.test_id = test_sessions.test_id").
			Where("sessions_selenium.session_id = ?", sessId).
			Where("sessions_selenium.valid = ?", true).
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
var ktestCache map[StringInt]KSessCacheEntry = make(map[StringInt]KSessCacheEntry)

// whether a key is valid for a test or not
func KeyIsValidForTest(key string, testId int64) bool {
	if key == "" {
		return false
	}
	entry, ok := ktestCache[StringInt{key, testId}]

	if !ok || time.Now().Unix()-entry.Time > int64(config.EnvConfig.CacheTTL) {
		// refresh cache
		var result bool
		err := config.DB.Model(&models.TestSession{}).
			Select("COUNT(*) > 0").
			Where("test_id = ?", testId).
			Where("Key = ?", key).
			Where("Current = ?", true).
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
		Time:  time.Now().Unix(),
		Valid: true,
	}
	ksessCache[StringPair{key, sessId}] = cacheEntry
	entry := models.SessionSelenium{
		Time:      cacheEntry.Time,
		TestId:    testId,
		SessionId: sessId,
		Valid:     true,
	}
	err := config.DB.Create(entry).Error
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
	err := config.DB.Model(&models.SessionSelenium{}).
		Where("session_id = ?", sessId).
		Update("valid", false).
		Error
	if err != nil {
		return err
	}
	return nil
}
