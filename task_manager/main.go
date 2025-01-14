package main

import (
	"task_manager/config"
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectToDatabase()

	r := gin.Default()

	r.POST("/create_task", controllers.CreateTask)
	r.GET("/get_tasks", controllers.GetTasks)
	r.GET("/task_details/:id", controllers.GetTaskDetails)

	r.Run(":8090")
}