package main

import (
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

const ENV_DB = "DB"
const ENV_DB_DEF = "postgres://thex:thex1234@db/thex"

func main() {
	dbStr := os.Getenv(ENV_DB)
	if dbStr == "" {
		dbStr = ENV_DB_DEF
	}
	db := DBInit(dbStr)
	err := db.AutoMigrate(&ApiKey{})
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
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
	err = db.AutoMigrate(&Session{})
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	err = db.AutoMigrate(&TestSession{})
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	log.Printf("AutoMigrate done")


	bytes, err := bcrypt.GenerateFromPassword([]byte("abcd1234"),
		bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed hashing password: %s", err.Error())
	}
	user := User{Username: "nafees", Password: string(bytes)}
	err = db.Create(&user).Error
	if err != nil {
		log.Fatalf("Failed to add user: %s", err.Error())
	}
}
