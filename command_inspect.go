package main

import "fmt"

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("a pokemon name is required")
	}

	name := args[0]

	pokemon, ok := cfg.caughtPokemon[name]

	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")

	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")

	for _, stat := range pokemon.Types {
		fmt.Printf("  - %s\n", stat.Type.Name)
	}
	return nil
}
