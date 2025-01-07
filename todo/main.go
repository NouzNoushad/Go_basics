package main

import (
	"gobasics/todo/config"
	"gobasics/todo/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectToDatabase()

	r := gin.Default()

	r.POST("/create_todo", controllers.CreateTodo)

	r.Run(":8000")
}
