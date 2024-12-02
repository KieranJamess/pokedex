package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

// commandCatch Attempt to catch a Pokemon.
func commandCatch(commands map[string]cliCommand, cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: catch <pokemon>")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	// The user has to roll a number 75% worthy of the Pokemon experience
	const percentile = 75
	catchCriteria := int(math.Ceil(float64(pokemon.BaseExperience) * percentile))
	numberToCatch := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Trying to catch %s. Pokemon experience: %v\n", pokemon.Name, pokemon.BaseExperience)
	if numberToCatch > catchCriteria {
		return fmt.Errorf("failed to catch %s", pokemon.Name)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("Caught %s!\n", pokemon.Name)

	return nil
}
