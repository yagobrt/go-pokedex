package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/yagobrt/go-pokedex/internal/pokeapi"
)

type config struct {
	commands      map[string]cliCommand
	pokeapiClient pokeapi.Client
	nextUrl       *string
	prevUrl       *string
	exploringArea *string
}

// Splits user input into words based on whitespace, trimming whitespace around
// words and returning a list of lowercase words
func cleanInput(text string) []string {
	lowercase := strings.ToLower(text)
	words := strings.Fields(lowercase)
	return words
}

// Run the REPL
func runREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]

		if commandName == "explore" {
			if len(input) < 2 {
				fmt.Println("Please provide an area to explore.")
				continue
			}
			cfg.exploringArea = &input[1]
		}
		cmd, ok := cfg.commands[commandName]
		if ok {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}
	}
}
