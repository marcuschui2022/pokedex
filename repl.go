package main

import (
	"bufio"
	"fmt"
	"github.com/marcuschui2022/pokedex/internal/pokecache"
	"os"
	"strings"
	"time"
)

type config struct {
	nextLocationsURL *string
	prevLocationsURL *string
	cache            *pokecache.Cache
}

func startRepl() {
	const cacheInterval = 5 * time.Second
	cache := pokecache.NewCache(cacheInterval)
	cfg := &config{
		cache: cache,
	}

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
				err := cmd.callback(cfg)
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
	callback    func(cfg *config) error
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
		"map": {
			name:        "map",
			description: "Displays the Pokedex map or the next page of the Pokedex map",
			callback:    commandMapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of the Pokedex map",
			callback:    commandMapBack,
		},
	}
}
