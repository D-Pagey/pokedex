package main

import (
	"time"

	"github.com/d-pagey/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.RespPokemon{},
	}

	startRepl(cfg)
}
