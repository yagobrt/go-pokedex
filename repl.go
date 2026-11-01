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
	caughtPokemon map[string]pokeapi.Pokemon
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

		args := []string{}
		if len(input) >= 2 {
			args = input[1:]
		}

		cmd, ok := cfg.commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := cmd.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}
