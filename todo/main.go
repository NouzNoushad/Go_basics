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
	r.GET("/get_todos", controllers.GetTodos)
	r.GET("/get_todo/:id", controllers.GetTodo)
    r.PUT("/update_todo/:id", controllers.UpdateTodo)
    r.DELETE("/delete_todo/:id", controllers.DeleteTodo)

	r.Run(":8000")
}
