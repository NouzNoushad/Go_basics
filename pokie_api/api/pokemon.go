package api

import (
	"fmt"
	"pokie_api/config"
	"pokie_api/models"
)

// Single pokemon by name or Id
func Pokemon(id string) (result models.Pokemon, err error) {
	err = config.Call(fmt.Sprintf("pokemon/%s", id), &result)
	return result, err
}