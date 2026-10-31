package main

import (
	"fmt"
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

	res, err := cfg.pokeapiClient.GetPokemon(*cfg.exploringArea)
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
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
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
	}
}
