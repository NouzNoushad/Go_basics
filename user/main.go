package main

import (
	"user/config"
	"user/controllers"
	"user/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectToDatabase()

	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	r.GET("/home", middleware.Authenticate(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Home page",
		})
	})

	r.Run(":8080")
}
