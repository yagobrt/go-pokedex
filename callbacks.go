package main

import (
	"fmt"
	"os"
)

// commandExit finishes the REPL loop
func commandExit() error {
	os.Exit(0)
	return nil
}

// commandHelp shows a help message and a list of availible commands
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
