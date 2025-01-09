package main

import (
	"user/config"
	"user/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectToDatabase()

	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to GO",
		})
	})

	r.Run(":8080")
}
