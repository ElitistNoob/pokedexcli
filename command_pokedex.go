package main

import "fmt"

func commandPokedex(cfg *config, _ *string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex.pokemons {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	fmt.Printf("You have %d Pokemons\n", cfg.pokedex.count)
	fmt.Println()
	fmt.Println("- - - - - - - -")

	return nil
}
