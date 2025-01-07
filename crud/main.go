package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	InSale bool    `json:"in_sale"`
}

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=noushad dbname=demo port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	} else {
		fmt.Println("Connected to database")
	}

	db.AutoMigrate(&Product{})
}

func main() {
	initDB()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to gin"})
	})

	r.Run(":8080")
}
