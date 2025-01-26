package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type config struct {
	Next     *string
	Previous *string
}

func startRepl() {
	cfg := &config{}
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
			description: "Displays a map of the Pokedex",
			callback:    commandMap,
		},
	}
}

type location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *config) error {
	var curURL string
	if cfg.Next != nil {
		curURL = *cfg.Next
	} else {
		curURL = "https://pokeapi.co/api/v2/location-area"
	}

	res, err := http.Get(curURL)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > http.StatusOK {
		return fmt.Errorf("HTTP status code %d", res.StatusCode)
	}

	var locations location
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return fmt.Errorf("%w", err)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
