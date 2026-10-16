package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Splits user input into words based on whitespace, trimming whitespace around
// words and returning a list of lowercase words
func cleanInput(text string) []string {
	lowercase := strings.ToLower(text)
	words := strings.Fields(lowercase)
	return words
}

// Run the REPL
func runREPL() {
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
		cmd, ok := getCommands()[commandName]
		if ok {
			cmd.callback()
		} else {
			fmt.Println("Unknown command")
		}
	}
}
