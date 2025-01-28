package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokemon struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(cfg *config) error {
	fmt.Printf("Exploring %s ...\n", *cfg.exploreArea)

	curURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", *cfg.exploreArea)

	data, err := httpGETExplore(curURL, cfg)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon")
	for _, poke := range data.PokemonEncounters {
		fmt.Println(" - " + poke.Pokemon.Name)
	}
	//fmt.Println()

	return nil
}

func httpGETExplore(url string, cfg *config) (pokemon, error) {
	// check cache
	if cachedData, found := cfg.cache.Get(url); found {
		//fmt.Println("Cache hit!")
		var pokemons pokemon
		err := json.Unmarshal(cachedData, &pokemons)
		if err != nil {
			return pokemon{}, err
		}
		return pokemons, nil
	}

	//fmt.Println("Cache miss! Fetching from API")
	res, err := http.Get(url)
	if err != nil {
		return pokemon{}, fmt.Errorf("failed to perform HTTP GET request to %s: %w", url, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error closing response body: %v", err)
		}
	}(res.Body)

	if res.StatusCode > http.StatusOK {
		return pokemon{}, fmt.Errorf("received unexpected HTTP status code %d from %s", res.StatusCode, url)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemon{}, fmt.Errorf("failed to read response body: %w", err)
	}

	//fmt.Println("Cache add!")
	cfg.cache.Add(url, body)

	var pokemons pokemon

	if err := json.Unmarshal(body, &pokemons); err != nil {
		return pokemon{}, fmt.Errorf("failed to decode response body: %w", err)
	} else {
		return pokemons, nil
	}
}
