package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// Exit the REPL
func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// Get help on commands
func commandHelp(cfg *config) error {
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
func commandMap(cfg *config) error {
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
func commandMapb(cfg *config) error {
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
func commandExplore(cfg *config) error {
	res, err := cfg.pokeapiClient.GetPokemon(*cfg.exploringArea)
	if err != nil {
		return err
	}

	for _, p := range res.PokemonEncounters {
		fmt.Println(p.Pokemon.Name)
	}
	return nil
}

// Generate map of all availible commands
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the names of 20 areas in the Pokemon world, each subsequent call will display the next 20 areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of the 20 previous areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display the names of the pokemon in one of the areas",
			callback:    commandExplore,
		},
	}
}
