package main

import (
	"time"

	"github.com/ElitistNoob/pokedexcli/api"
)

func main() {
	client := api.NewRequest(5 * time.Second)
	cfg := &config{
		apiClient: client,
	}
	startRepl(cfg)
}
