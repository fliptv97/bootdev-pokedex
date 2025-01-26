package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, params ...string) error {
	if len(params) < 1 {
		return errors.New("you must provide pokemon name")
	}

	name := params[0]
	pokemon, ok := cfg.pokedex[name]

	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  %v\n", t.Type.Name)
	}

	return nil
}
