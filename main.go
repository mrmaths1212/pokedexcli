package main

import (
	"time"

	"github.com/mrmaths1212/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
