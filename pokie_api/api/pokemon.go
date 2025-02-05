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

// Ability
func Ability(id string) (result models.Ability, err error) {
	err = config.Call(fmt.Sprintf("ability/%s", id), &result)
	return
}

// Pokemon form
func PokemonForm(id string) (result models.PokemonForm, err error) {
	err = config.Call(fmt.Sprintf("pokemon-form/%s", id), &result)
	return
}

// Characteristics
func Characteristics(id string) (result models.Characteristics, err error) {
	err = config.Call(fmt.Sprintf("characteristic/%s", id), &result)
	return
}

// Egg Groups
func EggGroups(id string) (result models.EggGroups, err error) {
	err = config.Call(fmt.Sprintf("egg-group/%s", id), &result)
	return result, err
}

// Gender
func Gender(id string) (result models.Gender, err error) {
	err = config.Call(fmt.Sprintf("gender/%s", id), &result)
	return
}

// Growth rate
func GrowthRate(id string) (result models.GrowthRates, err error) {
	err = config.Call(fmt.Sprintf("growth-rate/%s", id), &result)
	return
}

// Nature
func Nature(id string) (result models.Nature, err error) {
	err = config.Call(fmt.Sprintf("nature/%s", id), &result)
	return
}

// Pokemon Color
func PokemonColor(id string) (result models.PokemonColors, err error) {
	err = config.Call(fmt.Sprintf("pokemon-color/%s", id), &result)
	return result, err
}

// Pokemon Habitat
func PokemonHabitat(id string) (result models.PokemonHabitat, err error) {
	err = config.Call(fmt.Sprintf("pokemon-habitat/%s", id), &result)
	return
}

// Pokemon Shape
func PokemonShape(id string) (result models.PokemonShapes, err error) {
	err = config.Call(fmt.Sprintf("pokemon-shape/%s", id), &result)
	return
}

// Pokemon Species
func PokemonSpecies(id string) (result models.PokemonSpecies, err error) {
	err = config.Call(fmt.Sprintf("pokemon-species/%s", id), &result)
	return
}

// Stat
func Stat(id string) (result models.Stats, err error) {
	err = config.Call(fmt.Sprintf("stat/%s", id), &result)
	return
}

// Type
func Type(id string) (result models.Types, err error) {
	err = config.Call(fmt.Sprintf("type/%s", id), &result)
	return
}