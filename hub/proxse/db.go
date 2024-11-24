package main

import (
	"log"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// database
var db *gorm.DB

// initializes DB
func DBInit(connectionStr string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}
	db.AutoMigrate(&ApiKey{})
	db.AutoMigrate(&EventLogEntry{})
	return db
}

// API Key cache entry
type KCacheEntry struct {
	time  int64
	valid bool
}

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
