package main

import (
	"img_upload/config"
	"img_upload/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectToDatabase()

	r := gin.Default()

	r.POST("/upload_image", controllers.UploadImage)
	r.GET("/get_uploads", controllers.GetUploads)
	r.GET("/get_upload/:id", controllers.GetUploadById)

	r.Run(":8010")
}
