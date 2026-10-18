package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type PokemonResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetPokemon(area string) (PokemonResponse, error) {
	url := BaseURL + "/location-area/" + area

	// First check if the result is already in cache
	if val, ok := c.cache.Get(url); ok {
		pokemon := PokemonResponse{}
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return PokemonResponse{}, err
		}
		return PokemonResponse{}, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}

	pokemon := PokemonResponse{}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return PokemonResponse{}, err
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d, and\n body: %s\n", res.StatusCode, body)
	}

	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return PokemonResponse{}, err
	}

	c.cache.Add(url, body)
	return pokemon, nil
}
