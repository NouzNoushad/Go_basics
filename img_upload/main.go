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

	r.Run(":8010")
}
