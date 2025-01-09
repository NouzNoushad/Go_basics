package config

import (
	"fmt"
	"user/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectToDatabase() {
	dsn := "host=localhost user=postgres password=noushad dbname=user port=5432 sslmode=disable"

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	} else {
		fmt.Println("Connected to database")
	}

	db.AutoMigrate(&models.User{})
	Db = db
}
