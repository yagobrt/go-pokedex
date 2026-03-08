package main

import (
	"errors"
	"fmt"
	"os"
)

// cliCommand represents a command for the pokedex, with a name, a description
// and a callback function
type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

// getCommands returns a map with all the availible commands
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the names of 20 location areas in the Pokemon world. Each subsequent call will display the next 20 areas.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the name of the previous 20 location areas.",
			callback:    commandMapB,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s:\t%s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *config) error {
	res, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = res.Next
	cfg.prevLocationURL = res.Previous

	for _, l := range res.Locations {
		fmt.Println(l.Name)
	}
	return nil
}

func commandMapB(cfg *config) error {
	if cfg.prevLocationURL == nil {
		return errors.New("There are no previous locations.")
	}

	res, err := cfg.pokeapiClient.GetLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = res.Next
	cfg.prevLocationURL = res.Previous

	for _, l := range res.Locations {
		fmt.Println(l.Name)
	}
	return nil
}
