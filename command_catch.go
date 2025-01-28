package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemonResp, err := cfg.apiClient.CatchPokemon(name)
	if err != nil {
		return errors.New("invalid pokemon name")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	catchChance := float64(pokemonResp.BaseExperience) / 255.0 * 100
	randomChance := rand.Float64() * 100
	if randomChance > catchChance {
		fmt.Printf("%s escaped\n", name)
		return nil
	}

	cfg.caughtPokemon[name] = pokemonResp
	fmt.Printf("%s was caught!\n", name)

	return nil
}
