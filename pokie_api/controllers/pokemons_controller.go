package controllers

import (
	"net/http"
	"pokie_api/api"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPokemons(c *gin.Context) {
	endpoint := c.Param("name")
	offset := c.DefaultQuery("offset", "0")
	limit := c.DefaultQuery("limit", "10000")

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset number"})
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
		return
	}

	result, err := api.Pokemons(endpoint, offsetInt, limitInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// success
	c.JSON(http.StatusOK, gin.H{"data": result})
}