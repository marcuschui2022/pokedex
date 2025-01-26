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

func httpGET(url string, cfg *config) (location, error) {
	// check cache
	if cachedData, found := cfg.cache.Get(url); found {
		fmt.Println("Cache hit!")
		var locations location
		err := json.Unmarshal(cachedData, &locations)
		if err != nil {
			return location{}, err
		}
		return locations, nil
	}

	fmt.Println("Cache miss! Fetching from API")
	res, err := http.Get(url)
	if err != nil {
		return location{}, fmt.Errorf("failed to perform HTTP GET request to %s: %w", url, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error closing response body: %v", err)
		}
	}(res.Body)

	if res.StatusCode > http.StatusOK {
		return location{}, fmt.Errorf("received unexpected HTTP status code %d from %s", res.StatusCode, url)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return location{}, fmt.Errorf("failed to read response body: %w", err)
	}

	fmt.Println("Cache add!")
	cfg.cache.Add(url, body)

	var locations location

	if err := json.Unmarshal(body, &locations); err != nil {
		return location{}, fmt.Errorf("failed to decode response body: %w", err)
	} else {
		return locations, nil
	}
}

func commandMapForward(cfg *config) error {
	var curURL string
	if cfg.nextLocationsURL != nil {
		curURL = *cfg.nextLocationsURL
	} else {
		curURL = "https://pokeapi.co/api/v2/location-area"
	}

	data, err := httpGET(curURL, cfg)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = data.Next
	cfg.prevLocationsURL = data.Previous

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(cfg *config) error {
	var curURL string
	if cfg.prevLocationsURL != nil {
		curURL = *cfg.prevLocationsURL
	} else {

		return fmt.Errorf("you're on the first page")
	}

	data, err := httpGET(curURL, cfg)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = data.Next
	cfg.prevLocationsURL = data.Previous

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	return nil
}
