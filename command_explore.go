package main

import (
	"fmt"
)

func commandExplore(cfg *config, area *string) error {
	res, err := cfg.apiClient.ExploreArea(*area)
	if err != nil {
		return err
	}

	fmt.Printf("%d Pokemons found in %s:\n", res.Count, *area)
	for _, p := range res.Results {
		fmt.Printf(" - %s\n", p.Name)
	}
	fmt.Println()
	fmt.Println("- - - - - - - -")

	return nil
}
