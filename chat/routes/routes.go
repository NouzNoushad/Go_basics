package routes

import (
	"chat/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.POST("/register-user", handlers.RegisterUser)
	router.POST("/login-user", handlers.LoginUser)
	router.GET("/ws", handlers.HandleConnections)
}
