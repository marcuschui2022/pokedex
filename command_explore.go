package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	locationResp, err := cfg.apiClient.ListLocationPokemon(name)
	if err != nil {
		return errors.New("invalid location name")
	}
	fmt.Printf("Exploring %s...\n", locationResp.Name)
	fmt.Println("Found Pokemon")
	for _, data := range locationResp.PokemonEncounters {
		fmt.Println(" - " + data.Pokemon.Name)
	}
	//fmt.Println()

	return nil
}
