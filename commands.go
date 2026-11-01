package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

// Exit the REPL
func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// Get help on commands
func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range cfg.commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

// Get the next 20 locations
func commandMap(cfg *config, args ...string) error {
	res, err := cfg.pokeapiClient.GetLocations(cfg.nextUrl)
	if err != nil {
		return err
	}

	cfg.nextUrl = res.Next
	cfg.prevUrl = res.Previous

	for _, l := range res.Locations {
		fmt.Println(l.Name)
	}
	return nil
}

// Get the previous 20 locations
func commandMapB(cfg *config, args ...string) error {
	res, err := cfg.pokeapiClient.GetLocations(cfg.prevUrl)
	if err != nil {
		return err
	}

	cfg.nextUrl = res.Next
	cfg.prevUrl = res.Previous

	for _, l := range res.Locations {
		fmt.Println(l.Name)
	}
	return nil
}

// Get the pokemon in the area
func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Println("Please provide an area to explore.")
		return nil
	}
	cfg.exploringArea = &args[0]

	res, err := cfg.pokeapiClient.GetPokemonInArea(*cfg.exploringArea)
	if err != nil {
		return err
	}

	for _, p := range res.PokemonEncounters {
		fmt.Println(p.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Println("Please provide a pokemon to catch.")
		return nil
	}
	name := args[0]
	pokemonInfo, err := cfg.pokeapiClient.GetPokemonInfo(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	// Habría que hacer una segunda petición para obtener el ratio de captura
	// de la especie desde pokemonInfo.Species.URL para hacerlo bien
	chance := rand.Intn(pokemonInfo.BaseExperience)
	if chance > 50 {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)

	cfg.caughtPokemon[pokemonInfo.Name] = pokemonInfo
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Println("Please provide a pokemon to show info.")
		return nil
	}
	name := args[0]
	pokemonInfo, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("You have not caught that pokemon.")
	}

	fmt.Printf("Name: %s\n", pokemonInfo.Name)
	fmt.Printf("Height: %d\n", pokemonInfo.Height)
	fmt.Printf("Weight: %d\n", pokemonInfo.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf("\t- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	if len(pokemonInfo.Types) == 1 {
		fmt.Println("Type:")
		fmt.Printf("\t- %s\n", pokemonInfo.Types[0].Type.Name)
	} else {
		fmt.Println("Types:")
		for _, t := range pokemonInfo.Types {
			fmt.Printf("\t- %s\n", t.Type.Name)
		}
	}
	return nil
}

// Generate map of all availible commands
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays this help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the names of 20 areas in the Pokemon world, each subsequent call will display the next 20 areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of the 20 previous areas",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Display the names of the pokemon in one of the areas",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Show details of any catched pokemon",
			callback:    commandInspect,
		},
	}
}
