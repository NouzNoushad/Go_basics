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
	r.GET("/get-pokemon/:id", controllers.GetPokemon)
	r.GET("/get-ability/:id", controllers.GetPokemonAbility)
	r.GET("/get-pokemon-form/:id", controllers.GetPokemonForm)
	r.GET("/get-characteristic/:id", controllers.GetCharacteristics)
	r.GET("/get-egg-group/:id", controllers.GetEggGroups)
	r.GET("/get-gender/:id", controllers.GetGenders)
	r.GET("/get-growth-rate/:id", controllers.GetGrowthRates)
	r.GET("/get-nature/:id", controllers.GetNature)
	r.GET("/get-pokemon-color/:id", controllers.GetPokemonColors)
	r.GET("/get-pokemon-habitat/:id", controllers.GetPokemonHabitat)
	r.GET("/get-pokemon-shape/:id", controllers.GetPokemonShapes)
	r.GET("/get-pokemon-species/:id", controllers.GetPokemonSpecies)
	r.GET("/get-stat/:id", controllers.GetStat)
	r.GET("/get-type/:id", controllers.GetTypes)

	r.GET("/get-pokemons/:name", controllers.GetPokemons)

	r.Run(":8012")
}
