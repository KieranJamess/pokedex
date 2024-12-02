package main

import "fmt"

// commandHelp displays the help message.
func commandHelp(commands map[string]cliCommand, cfg *config) error {
	fmt.Println("Available commands:")
	for _, cmd := range commands {
		fmt.Printf("%s - %s\n", cmd.name, cmd.description)
	}
	return nil
}
