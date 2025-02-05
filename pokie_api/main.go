package main

import (
	"pokie_api/config"
	"pokie_api/controllers"

	"github.com/gin-gonic/gin"
)

// pokemon apis
func main() {
	config.Init()

	r := gin.Default()
	r.GET("/get_pokemon/:id", controllers.GetPokemon)

	r.Run(":8012")
}
