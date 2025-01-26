package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, params ...string) error {
	if len(params) < 1 {
		return errors.New("you must provide pokemon name")
	}

	name := params[0]

	fmt.Printf("Throwing a Pokeball at %v...\n", name)

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)

	if err != nil {
		return err
	}

	if rand.IntN(pokemon.BaseExperience) <= 40 {
		fmt.Printf("%v was caught!\n", name)

		cfg.pokedex[name] = pokemon
	} else {
		fmt.Printf("%v escaped!\n", name)
	}

	return nil
}
