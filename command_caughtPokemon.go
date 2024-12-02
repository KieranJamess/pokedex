package main

import (
	"fmt"
)

// commandCaughtPokemon displays all caught Pokemon.
func commandCaughtPokemon(commands map[string]cliCommand, cfg *config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("caught pokemon doesn't accept any arguments")
	}

	if len(cfg.caughtPokemon) == 0 {
		return fmt.Errorf("you haven't caught any pokemon yet")
	}

	fmt.Println("You have caught following pokemon:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}
