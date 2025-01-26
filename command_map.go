package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func commandMapForward(cfg *config) error {
	var curURL string
	if cfg.nextLocationsURL != nil {
		curURL = *cfg.nextLocationsURL
	} else {
		curURL = "https://pokeapi.co/api/v2/location-area"
	}

	res, err := http.Get(curURL)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error closing response body: %v", err)
		}
	}(res.Body)

	if res.StatusCode > http.StatusOK {
		return fmt.Errorf("HTTP status code %d", res.StatusCode)
	}

	var locations location
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return fmt.Errorf("%w", err)
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(cfg *config) error {
	var curURL string
	if cfg.prevLocationsURL != nil {
		curURL = *cfg.prevLocationsURL
	} else {

		return fmt.Errorf("no previous page")
	}

	res, err := http.Get(curURL)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error closing response body: %v", err)
		}
	}(res.Body)

	if res.StatusCode > http.StatusOK {
		return fmt.Errorf("HTTP status code %d", res.StatusCode)
	}

	var locations location
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return fmt.Errorf("%w", err)
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
