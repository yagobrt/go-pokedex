package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

func runREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from std in")
		}
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		cmd, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := cmd.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}
