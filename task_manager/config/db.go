package config

import (
	"fmt"
	"task_manager/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var TaskDB *gorm.DB
var UserDB *gorm.DB

func ConnectToDatabase() {
	// task manager
	dsnTaskManager := "host=localhost user=postgres password=noushad dbname=task-manager port=5432 sslmode=disable"

	var err error
	dbTaskManager, err := gorm.Open(postgres.Open(dsnTaskManager), &gorm.Config{})
	if err != nil {
		panic("Failed to connect task manager database")
	} else {
		fmt.Println("Connected to task manager database")
	}

	// user manager
	dsnUserManager := "host=localhost user=postgres password=noushad dbname=user-manager port=5432 sslmode=disable"
	
	dbUserManager, err := gorm.Open(postgres.Open(dsnUserManager), &gorm.Config{})
	if err != nil {
		panic("Failed to connect user manager database")
	} else {
		fmt.Println("Connected to user manager database")
	}

	dbTaskManager.AutoMigrate(&models.Task{})
	dbUserManager.AutoMigrate(&models.User{})

	TaskDB = dbTaskManager
	UserDB = dbUserManager
}