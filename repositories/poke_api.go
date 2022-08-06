package repositories

import (
	"errors"
	"fmt"

	"GoConcurrency-Bootcamp-2022/models"

	"github.com/go-resty/resty/v2"
)

type PokeAPI struct{}

const url = "https://pokeapi.co/api/v2/"

func (pa PokeAPI) FetchPokemon(id int) (models.Pokemon, error) {
	client := resty.New()

	poke := models.Pokemon{}

	_, err := client.
		R().
		SetHeader("Content-Type", "application/json").
		SetResult(&poke).
		Post(fmt.Sprintf("%s/pokemon/%d", url, id))
	if err != nil {
		fmt.Println("error on request!!")
		return models.Pokemon{}, err
	}
	if poke.Name == "" {
		nonExistingData := errors.New("Error obtaining pokemon data.")
		return models.Pokemon{}, nonExistingData
	}
	return poke, nil
}

func (pa PokeAPI) FetchAbility(url string) (models.Ability, error) {
	ability := models.Ability{}

	client := resty.New()

	_, err := client.
		R().
		SetHeader("Content-Type", "application/json").
		SetResult(&ability).
		Get(url)

	if err != nil {
		return models.Ability{}, err
	}

	return ability, nil
}
