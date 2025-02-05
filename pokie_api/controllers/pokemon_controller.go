package controllers

import (
	"net/http"
	"pokie_api/api"

	"github.com/gin-gonic/gin"
)

func GetPokemon(c *gin.Context) {
	id := c.Param("id")

	result, err := api.Pokemon(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// success
	c.JSON(http.StatusOK, gin.H{"pokemon": result})
}