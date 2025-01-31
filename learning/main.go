package main

import (
	"learning/config"
	"learning/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectToDatabase()

	r := gin.Default()

	r.POST("/create_learning", controllers.CreateLearning)
	r.POST("/create_chapter", controllers.CreateChapter)
	r.POST("/create_study_material", controllers.CreateStudyMaterial)

	r.Run(":8011")
}