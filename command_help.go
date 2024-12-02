package main

import "fmt"

// commandHelp displays the help message.
func commandHelp(commands map[string]cliCommand, cfg *config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("help doesn't accept any arguments")
	}
	fmt.Println("Available commands:")
	for _, cmd := range commands {
		fmt.Printf("%s - %s\n", cmd.name, cmd.description)
	}
	return nil
}
