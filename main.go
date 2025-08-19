package main

import (
	"time"

	"github.com/yagobrt/go-pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	runREPL(cfg)
}
