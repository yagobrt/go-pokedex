package main

// cliCommand represents a command for the pokedex, with a name, a description
// and a callback function
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// getCommands returns a map with all the availible commands
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex.",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the names of 20 location areas in the Pokemon world. Each subsequent call will display the next 20 areas.",
			callback:    func() error { return nil },
		},
		"mapb": {
			name:        "mapb",
			description: "Display the name of the previous 20 location areas.",
			callback:    func() error { return nil },
		},
	}

}
