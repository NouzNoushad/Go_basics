package main

import (
	"chat/database"
	"chat/handlers"
	"chat/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()

	routes.Routes(r)

	go handlers.HandleMessages()

	r.Run(":8015")
}