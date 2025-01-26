package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, params ...string) error {
	if len(params) < 1 {
		return errors.New("you should provide location area name")
	}

	locationAreaName := params[0]

	fmt.Printf("Exploring %v...\n", locationAreaName)

	locationAreaRes, err := cfg.pokeapiClient.GetLocationAreaByName(locationAreaName)

	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locationAreaRes.PokemonEncounters {
		fmt.Printf(" - %v\n", encounter.Pokemon.Name)
	}

	return nil
}
