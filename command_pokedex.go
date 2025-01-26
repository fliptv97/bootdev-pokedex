package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, _ ...string) error {
	if len(cfg.pokedex) == 0 {
		return errors.New("sorry, but your pokedex is empty")
	}

	fmt.Println("Your Pokedex:")
	for pokemonName := range cfg.pokedex {
		fmt.Println("  ", pokemonName)
	}

	return nil
}
