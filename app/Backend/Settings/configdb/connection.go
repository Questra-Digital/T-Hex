package configdb

import (
	"fmt"
	"log"
	"time"

	"Settings/graph/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectCockroachDB() *gorm.DB { //using function return instead
	dsn := "postgresql://waleedmahmood:H1YRCS6CPgQursPyvRx8LQ@metal-kakapo-3678.8nk.cockroachlabs.cloud:26257/basic-app?sslmode=verify-full"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	db.AutoMigrate(&model.Settings{})

	var now time.Time
	db.Raw("SELECT NOW()").Scan(&now)

	fmt.Println(now)
	return db
}
