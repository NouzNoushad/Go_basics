package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "status ok"})
	})

	return router
}
