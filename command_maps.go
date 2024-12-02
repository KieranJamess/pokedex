package main

import (
	"fmt"
)

// commandHelp displays the help message.
func commandMap(commands map[string]cliCommand, cfg *config) error {
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Available locations:")
	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = res.Next
	cfg.previousLocationAreaURL = res.Previous

	return nil
}

// commandHelp displays the help message.
func commandMapback(commands map[string]cliCommand, cfg *config) error {
	if cfg.previousLocationAreaURL == nil {
		return fmt.Errorf("no previous location area (you're on the first page!)")
	}

	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Available locations:")
	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = res.Next
	cfg.previousLocationAreaURL = res.Previous

	return nil
}
