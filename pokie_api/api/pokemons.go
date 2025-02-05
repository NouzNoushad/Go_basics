package api

import (
	"fmt"
	"pokie_api/config"
	"pokie_api/models"
)

func Pokemons(endpoint string, offset int, limit int) (result models.Pokemons, err error) {
	err = config.Call(fmt.Sprintf("%s?offset=%d&limit=%d", endpoint, offset, limit), &result)
	return
}