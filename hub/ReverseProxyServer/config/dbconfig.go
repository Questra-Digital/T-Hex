package config

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func DBInit() *gorm.DB {
	dbStr := EnvConfig.DbUrl
	if dbStr == "" {
		dbStr = "postgres://thex:thex1234@db/thex"
	}
	db, err := gorm.Open(postgres.Open(dbStr), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}
	DB = db
	return db
}

func DBClose() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to close database: %s", err.Error())
	}
	sqlDB.Close()
	log.Printf("Database closed successfully")
}