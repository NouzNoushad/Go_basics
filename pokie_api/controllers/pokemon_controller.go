package controllers

import (
	"net/http"
	"pokie_api/api"

	"github.com/gin-gonic/gin"
)

// Pokemon
func GetPokemon(c *gin.Context) {
	GetPokemonData(c, "pokemon", func(s string) (interface{}, error) {
		return api.Pokemon(s)
	})
}

// Pokemon ability
func GetPokemonAbility(c *gin.Context) {
	GetPokemonData(c, "ability", func(s string) (interface{}, error) {
		return api.Ability(s)
	})
}

// Pokemon form
func GetPokemonForm(c *gin.Context) {
	GetPokemonData(c, "pokemon form", func(s string) (interface{}, error) {
		return api.PokemonForm(s)
	})
}

// Characteristics
func GetCharacteristics(c *gin.Context) {
	GetPokemonData(c, "characteristics", func(s string) (interface{}, error) {
		return api.Characteristics(s)
	})
}

// Egg Groups
func GetEggGroups(c *gin.Context) {
	GetPokemonData(c, "egg groups", func(s string) (interface{}, error) {
		return api.EggGroups(s)
	})
}

// Genders
func GetGenders(c *gin.Context) {
	GetPokemonData(c, "genders", func(s string) (interface{}, error) {
		return api.Gender(s)
	})
}

// Growth rates
func GetGrowthRates(c *gin.Context) {
	GetPokemonData(c, "growth rate", func(s string) (interface{}, error) {
		return api.GrowthRate(s)
	})
}

// Nature
func GetNature(c *gin.Context) {
	GetPokemonData(c, "nature", func(s string) (interface{}, error) {
		return api.Nature(s)
	})
}

// Pokemon Colors
func GetPokemonColors(c *gin.Context) {
	GetPokemonData(c, "pokemon color", func(s string) (interface{}, error) {
		return api.PokemonColor(s)
	})
}

// Pokemon Habitat
func GetPokemonHabitat(c *gin.Context) {
	GetPokemonData(c, "pokemon habitat", func(s string) (interface{}, error) {
		return api.PokemonHabitat(s)
	})
}

// Pokemon Shapes
func GetPokemonShapes(c *gin.Context) {
	GetPokemonData(c, "pokemon shape", func(s string) (interface{}, error) {
		return api.PokemonShape(s)
	})
}

// Pokemon Species
func GetPokemonSpecies(c *gin.Context) {
	GetPokemonData(c, "pokemon species", func(s string) (interface{}, error) {
		return api.PokemonSpecies(s)
	})
}

// Stat
func GetStat(c *gin.Context) {
	GetPokemonData(c, "stat", func(s string) (interface{}, error) {
		return api.Stat(s)
	})
}

// Types
func GetTypes(c *gin.Context) {
	GetPokemonData(c, "type", func(s string) (interface{}, error) {
		return api.Type(s)
	})
}

// Get pokemon data
func GetPokemonData(c *gin.Context, responseType string, apiFunc func(string) (interface{}, error)) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid id"})
		return
	}

	result, err := apiFunc(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// success
	c.JSON(http.StatusOK, gin.H{responseType: result})
}
