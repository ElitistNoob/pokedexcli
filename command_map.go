package main

import (
	"fmt"
	"time"

	cache "github.com/ElitistNoob/pokedexcli/internal/pokecache"
)

var pkCache = cache.NewCache(5 * time.Minute)

func commandMap(cfg *config) error {
	res, err := cfg.apiClient.GetLocations(cfg.next, pkCache)
	if err != nil {
		return err
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

	res, err := cfg.apiClient.GetLocations(cfg.previous, pkCache)
	if err != nil {
		return err
	}

	cfg.next = res.Next
	cfg.previous = res.Previous

	for _, d := range res.Results {
		fmt.Println(d.Name)
	}

	return nil
}
