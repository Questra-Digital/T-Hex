package main

import (
	"os"
	"log"
)

const ENV_DB = "DB"
const ENV_DB_DEF = "postgres://thex:thex1234@db/thex"

func main() {
	dbStr := os.Getenv(ENV_DB)
	if dbStr == "" {
		dbStr = ENV_DB_DEF
	}
	DBInit(dbStr)
	if err := db.AutoMigrate(&ApiKey{}); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	if err := db.AutoMigrate(&UserKey{}); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	if err := db.AutoMigrate(&EventLogEntry{}); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	if err := db.AutoMigrate(&KeySession{}); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	log.Printf("AutoMigrate done")
}
