package main

import (
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	db := DBInit()
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
	err = db.AutoMigrate(&ContactUsMessage{})
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	log.Printf("AutoMigrate done")

	if os.Getenv("AM_DEMO") == "" {
		return
	}

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

	if err := db.Create(&ApiKey{Key: "abcd1234"}).Error; err != nil {
		log.Fatalf("Failed to add API Key: %s", err.Error())
	}

	if err := db.Create(&UserKey{"nafees", "abcd1234"}).Error; err != nil {
		log.Fatalf("Failed to add User-Key relation: %s", err.Error())
	}

}
