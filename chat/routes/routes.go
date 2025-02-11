package routes

import (
	"chat/handlers"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {

	router := gin.Default()

	router.POST("/register-user", handlers.RegisterUser)
	router.POST("/login-user", handlers.LoginUser)

	return router
}
