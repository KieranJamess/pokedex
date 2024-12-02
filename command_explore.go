package main

import (
	"errors"
	"fmt"
)

// commandExplore explore a area provided.
func commandExplore(commands map[string]cliCommand, cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: explore <location>")
	}

	areaName := args[0]

	res, err := cfg.pokeapiClient.GetLocationData(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon found in %s:\n", areaName)

	for _, pokemon := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
