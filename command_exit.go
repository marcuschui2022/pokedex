package main

import (
	"fmt"
	"os"
)

// commandExit terminates the program with a successful exit status and returns nil error.
func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
