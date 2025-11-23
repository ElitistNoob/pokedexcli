package main

import (
	"fmt"
)

func commandInspect(cfg *config, name *string) error {
	if name == nil {
		fmt.Println("no pokemon was provided")
		fmt.Println()
		fmt.Println("- - - - - - - -")
		return nil
	}

	pokemon := *name
	p, caught := cfg.pokedex.pokemons[pokemon]
	if !caught {
		fmt.Println("you have not caught that pokemon")
		fmt.Println()
		fmt.Println("- - - - - - - -")
		return nil
	}

	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Println("Stats:")
	for _, v := range p.Stats {
		fmt.Printf("  - %s: %v\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Println("Types:")
	for _, v := range p.Types {
		fmt.Printf("  - %s\n", v.Type.Name)
	}
	fmt.Println()
	fmt.Println("- - - - - - - -")
	return nil
}
