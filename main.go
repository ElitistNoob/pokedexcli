package main

import (
	"time"

	api "github.com/ElitistNoob/pokedexcli/api"
)

func main() {
	client := api.NewRequest(5*time.Second, 5*time.Minute)
	cfg := &config{
		apiClient: client,
		pokedex:   pokedex{pokemons: make(map[string]api.Pokemon), count: 0},
	}
	startRepl(cfg)
}
