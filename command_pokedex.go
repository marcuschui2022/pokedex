package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	_ = args
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You have not caught any Pokemon yet!")
		return nil
	}
	fmt.Println("Your Pokedex:")

	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
