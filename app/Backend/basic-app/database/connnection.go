package database

import (
	"basic-app/models"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB //global variable required by authController

func Connect() {
	dsn := "postgresql://waleedmahmood:H1YRCS6CPgQursPyvRx8LQ@metal-kakapo-3678.8nk.cockroachlabs.cloud:26257/basic-app?sslmode=verify-full"
	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	DB = dbConnection

	dbConnection.AutoMigrate(&models.User{})

	var now time.Time
	DB.Raw("SELECT NOW()").Scan(&now)

	fmt.Println(now)
}
