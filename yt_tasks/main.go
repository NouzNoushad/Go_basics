package main

import (
	"yt_tasks/config"
	"yt_tasks/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectToDatabase()

	r := gin.Default()

	r.Static("/uploads", "./uploads")

	r.POST("/create_task", controllers.CreateTask)
	r.GET("/get_tasks", controllers.GetTasks)
	r.GET("/task_details/:id", controllers.GetTaskDetails)
	r.PUT("/update_task_details/:id", controllers.UpdateTaskDetails)
	r.DELETE("/delete_task_details/:id", controllers.DeleteTaskDetails)

	r.Run(":8020")
}
