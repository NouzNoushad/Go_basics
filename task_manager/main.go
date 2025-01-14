package main

import (
	"net/http"
	"task_manager/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectToDatabase()
	
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to task manager"})
	})

	r.Run(":8090")
}