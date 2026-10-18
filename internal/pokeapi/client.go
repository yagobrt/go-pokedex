package pokeapi

import (
	"net/http"
	"time"

	"github.com/yagobrt/go-pokedex/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// Create a new client
func NewClient(cacheInterval, timeout time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
