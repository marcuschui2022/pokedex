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

		if cmd, cmdExists := commands[commandName]; cmdExists {
			if err := cmd.callback(); err != nil {
				fmt.Printf("Error executing command: %s\n", err)
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

// commandHelp displays a help message with descriptions of available commands and their usage.
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

// commandExit terminates the program with a successful exit status and returns nil error.
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// cliCommand represents a single CLI command with its name, description, and a callback function to execute it.
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// commands is a map of supported CLI commands, mapping command names to their corresponding cliCommand struct.
var commands = map[string]cliCommand{
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
