package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("a pokemon name is required")
	}

	name := args[0]

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	threshold := 40
	random := rand.Intn(pokemonResp.BaseExperience)

	if random <= threshold {
		fmt.Printf("%s was caught!\n", name)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.caughtPokemon[name] = pokemonResp
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
