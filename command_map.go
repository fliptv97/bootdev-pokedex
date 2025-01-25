package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	locationAreasRes, err := cfg.pokeapiClient.LoadLocationAreas(cfg.nextLocationAreasUrl)

	if err != nil {
		return err
	}

	cfg.prevLocationAreasUrl = locationAreasRes.Prev
	cfg.nextLocationAreasUrl = locationAreasRes.Next

	for _, locationArea := range locationAreasRes.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationAreasUrl == nil {
		return errors.New("you're on the first page")
	}

	locationAreasRes, err := cfg.pokeapiClient.LoadLocationAreas(cfg.prevLocationAreasUrl)

	if err != nil {
		return err
	}

	cfg.prevLocationAreasUrl = locationAreasRes.Prev
	cfg.nextLocationAreasUrl = locationAreasRes.Next

	for _, locationArea := range locationAreasRes.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}
