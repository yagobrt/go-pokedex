package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("pokedex> ")
		scanner.Scan()
		commandName := scanner.Text()
		// TODO: Clean input (lowercase, separate words...)

		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

		err := scanner.Err()
		if err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

	}
}
