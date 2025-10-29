package main

import (
	"fmt"
)

var BASE_ENDPOINT = "https://pokeapi.co/api/v2/location-area"

func commandMap(cfg *config) error {
	res, err := cfg.apiClient.GetLocations(cfg.next)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	cfg.next = res.Next
	cfg.previous = res.Previous

	for _, d := range res.Results {
		fmt.Println(d.Name)
	}

	return nil
}

func commandMapBack(cfg *config) error {
	if cfg.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := cfg.apiClient.GetLocations(cfg.previous)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	cfg.next = res.Next
	cfg.previous = res.Previous

	for _, d := range res.Results {
		fmt.Println(d.Name)
	}

	return nil
}
