package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, _ ...string) error {
	locationAreasRes, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationAreasUrl)

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

func commandMapb(cfg *config, _ ...string) error {
	if cfg.prevLocationAreasUrl == nil {
		return errors.New("you're on the first page")
	}

	locationAreasRes, err := cfg.pokeapiClient.GetLocationAreas(cfg.prevLocationAreasUrl)

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
