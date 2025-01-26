package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		if cmd, cmdExists := getCommands()[commandName]; cmdExists {
			if cmdExists {
				err := cmd.callback()
				if err != nil {
					fmt.Printf("Error executing command: %s\n", err)
				}
			}
		} else {
			fmt.Printf("Unknown command: %s\n", commandName)
		}

	}
}

// cleanInput processes a string by converting it to lowercase, splitting it into words, and returning a slice of words.
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

// cliCommand represents a single CLI command with its name, description, and a callback function to execute it.
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// getCommands returns a map of available CLI commands, each with its name, description, and associated callback function.
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
	}
}
