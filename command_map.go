package main

import (
	"fmt"
)

func commandMap(cfg *config, _ *string) error {
	res, err := cfg.apiClient.GetLocations(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = res.Next
	cfg.previous = res.Previous

	for _, d := range res.Results {
		fmt.Printf(" - %s\n", d.Name)
	}
	fmt.Println()
	fmt.Println("- - - - - - - -")

	return nil
}

func commandMapBack(cfg *config, _ *string) error {
	if cfg.previous == nil {
		fmt.Println("you're on the first page")
		fmt.Println()
		fmt.Println("- - - - - - - -")
		return nil
	}

	res, err := cfg.apiClient.GetLocations(cfg.previous)
	if err != nil {
		return err
	}

	cfg.next = res.Next
	cfg.previous = res.Previous

	for _, d := range res.Results {
		fmt.Printf(" - %s\n", d.Name)
	}
	fmt.Println()
	fmt.Println("- - - - - - - -")

	return nil
}
