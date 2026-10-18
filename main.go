package main

import (
	"time"

	"github.com/yagobrt/go-pokedex/internal/pokeapi"
)

func main() {
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeapi.NewClient(5*time.Minute, 5*time.Second),
		exploringArea: nil,
	}
	runREPL(cfg)
}
