package main

import "fmt"

// commandHelp displays a help message with descriptions of available commands and their usage.
func commandHelp(cfg *config, args ...string) error {
	_ = args
	_ = cfg
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()

	return nil
}
