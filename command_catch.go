package main

import "fmt"

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("a pokemon name is required")
	}

	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	return nil
}
