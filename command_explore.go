package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no args passed")
	}

	area := args[0]
	fmt.Println(area)
	return nil
}
