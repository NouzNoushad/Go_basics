package config

import (
	"fmt"
	"img_upload/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	dsn := "host=localhost user=postgres password=noushad dbname=upload port=5432 sslmode=disable"

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	} else {
		fmt.Println("Connected to database")
	}

	db.AutoMigrate(&models.Upload{})
	DB = db
}