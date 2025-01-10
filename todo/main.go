package main

import (
	"gobasics/config"
	"gobasics/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectToDatabase()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	r.POST("/create_todo", controllers.CreateTodo)
	r.GET("/get_todos", controllers.GetTodos)
	r.GET("/get_todo/:id", controllers.GetTodo)
	r.PUT("/update_todo/:id", controllers.UpdateTodo)
	r.DELETE("/delete_todo/:id", controllers.DeleteTodo)

	r.Run(":8000")
}
