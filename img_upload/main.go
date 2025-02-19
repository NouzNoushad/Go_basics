package main

import (
	"img_upload/config"
	"img_upload/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectToDatabase()

	r := gin.Default()

	r.Static("/uploads", "./uploads")

	r.POST("/upload_image", controllers.UploadImage)
	r.GET("/get_uploads", controllers.GetUploads)
	r.GET("/get_details/:id", controllers.GetUploadById)
	r.PUT("/update_details/:id", controllers.UpdateUpload)
	r.DELETE("/delete_details/:id", controllers.DeleteUpload)

	r.Run(":8010")
}
