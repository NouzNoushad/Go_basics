package config

import (
	"fmt"
	"log"
	"user_auth/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "noushad"
	dbname   = "auth_db"
	sslmode  = "disable"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	} else{
		fmt.Println("Connected to database")
	}

	// auto migrate
	db.AutoMigrate(&models.User{})

	DB = db
}