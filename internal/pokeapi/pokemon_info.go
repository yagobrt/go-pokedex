package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetPokemonInfo(pokemon string) (Pokemon, error) {
	url := BaseURL + "/pokemon/" + pokemon

	// Check cache
	if val, ok := c.cache.Get(url); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonInfo := Pokemon{}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return Pokemon{}, err
	}

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d, and\n body: %s\n", res.StatusCode, body)
	}

	err = json.Unmarshal(body, &pokemonInfo)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, body)
	return pokemonInfo, nil
}
