package main

import (
	"errors"
	"fmt"
)

// commandInspect inspect a Pokemon you have caught.
func commandInspect(commands map[string]cliCommand, cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: inspect <pokemon>")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("pokemon not caught")
	}

	fmt.Printf("- Name: %s\n", pokemon.Name)
	fmt.Printf("- Experience: %v\n", pokemon.BaseExperience)
	fmt.Printf("- Height: %v\n", pokemon.Height)
	fmt.Printf("- Weight: %v\n", pokemon.Weight)
	fmt.Printf("- Types:\n")
	for _, typ := range pokemon.Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}

	return nil
}
