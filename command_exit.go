package main

import (
	"fmt"
	"os"
)

// commandExit exits the application.
func commandExit(commands map[string]cliCommand, cfg *config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("exit doesn't accept any arguments")
	}
	fmt.Println("Exiting the application.")
	os.Exit(0)
	return nil
}
