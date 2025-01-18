package main

import (
	"net/http"
	"yt_tasks/config"
	"yt_tasks/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectToDatabase()

	r := gin.Default()

	r.POST("/create_task", controllers.CreateTask)
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to yt tasks"})
	})

	r.Run(":8020")
}
