package main

import (
	"learning/config"
	"learning/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectToDatabase()

	r := gin.Default()

	r.Static("/uploads", "./uploads")

	r.POST("/create_learning", controllers.CreateLearning)
	r.POST("/create_chapter", controllers.CreateChapter)
	r.POST("/create_study_material", controllers.CreateStudyMaterial)
	r.GET("/get_learning", controllers.GetLearningList)
	r.GET("/get_learning_details/:id", controllers.GetLearningById)
	r.DELETE("/delete_learning_details/:id", controllers.DeleteLearning)
	r.PUT("/update_study_material/:id", controllers.CreateStudyMaterial)
	r.PUT("/update_chapter/:id", controllers.UpdateChapter)
	r.PUT("/update_learning/:id", controllers.UpdateLearning)

	r.Run(":8011")
}
