package main

import (
	"strings"
)

func cleanInput(text string) []string {
	trimmed := strings.Trim(text, " ")
	words := strings.Split(trimmed, " ")
	lowered := []string{}

	for _, w := range words {
		lowered = append(lowered, strings.ToLower(w))
	}

	return lowered
}
