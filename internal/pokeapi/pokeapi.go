package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/yagobrt/go-pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	pokecache.Cache
}

func NewClient(timeout, cacheDuration time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		Cache: pokecache.NewCache(cacheDuration),
	}
}

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
	url := "https://pokeapi.co/api/v2/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	// Read from cache if we already have it
	if val, ok := c.Cache.Get(url); ok {
		loc := LocationResponse{}
		err := json.Unmarshal(val, &loc)
		if err != nil {
			return LocationResponse{}, err
		}
		return loc, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	loc := LocationResponse{}
	err = json.Unmarshal(body, &loc)
	if err != nil {
		return LocationResponse{}, err
	}

	c.Cache.Add(url, body)
	return loc, nil
}
