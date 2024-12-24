package main

import (
	"log"
	"os"
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
}
