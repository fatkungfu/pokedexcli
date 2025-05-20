package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("Your Pokedex is empty!")
	}

	fmt.Println("Your Pokedex:")
	for _, p := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n", p.Name)
	}
	return nil
}
