package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const BaseURL = "https://pokeapi.co/api/v2"

type LocationResponse struct {
	Count     int     `json:"count"`
	Next      *string `json:"next"`
	Previous  *string `json:"previous"`
	Locations []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocations(pageUrl *string) (LocationResponse, error) {
	url := BaseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	// First check if the result is already in cache
	if val, ok := c.cache.Get(url); ok {
		locations := LocationResponse{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return LocationResponse{}, err
		}
		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResponse{}, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return LocationResponse{}, err
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d, and\n body: %s\n", res.StatusCode, body)
	}

	locations := LocationResponse{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return LocationResponse{}, err
	}

	c.cache.Add(url, body)
	return locations, nil
}
