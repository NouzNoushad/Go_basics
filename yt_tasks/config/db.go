package config

import (
	"fmt"
	"yt_tasks/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	dsn := "host=localhost user=postgres password=noushad dbname=manager port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	} else {
		fmt.Println("Connected to database")
	}

	db.AutoMigrate(&models.Task{})

	DB = db
}