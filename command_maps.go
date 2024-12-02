package main

import (
	"fmt"
)

// commandMap shows the first 20 map entries. Keep using to iterate through.
func commandMap(commands map[string]cliCommand, cfg *config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("map doesn't accept any arguments")
	}
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

// commandMapback Move back a page from the maps.
func commandMapback(commands map[string]cliCommand, cfg *config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("mapback doesn't accept any arguments")
	}
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
