package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func successChance(baseExp int) float64 {
	x := float64(baseExp)
	chance := 100 - math.Log(x)*15
	if chance < 25 {
		chance = 25
	}
	return chance
}

func attempt(baseExp int, r *rand.Rand) bool {
	threshold := successChance(baseExp)
	roll := r.Float64() * 100

	return roll < threshold
}

func commandCatch(cfg *config, arg *string) error {
	if arg == nil {
		fmt.Println("a pokemon name must be provided")
		return nil
	}

	n := *arg
	fmt.Printf("Throwing a Pokeball at %s...\n", n)

	pokemon, err := cfg.apiClient.GetPokemon(n)
	if err != nil {
		return err
	}

	pokedex := cfg.pokedex.pokemons
	if _, exist := pokedex[n]; !exist {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		ok := attempt(pokemon.BaseExperience, r)
		if !ok {
			fmt.Printf("%s Escaped!\n", n)
			fmt.Println()
			fmt.Println("- - - - - - - -")
			return nil
		}
		fmt.Printf("%s Caught!\n", n)
		pokedex[n] = *pokemon
		cfg.pokedex.count++
	}
	fmt.Println()
	fmt.Println("- - - - - - - -")

	return nil
}
