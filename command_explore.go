package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no args passed")
	}

	area := args[0]

	fmt.Println("Exploring " + area + "...")

	locationResp, err := cfg.pokeapiClient.GetLocation(area)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationResp.PokemonEncounters {
		fmt.Println(" - " + pokemon.Pokemon.Name)
	}

	return nil
}
