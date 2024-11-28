package main

import (
	"log"
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
	db.AutoMigrate(&KeySession{})
	return db
}
