package main

import (
	"fmt"
	"os"
)

// commandExit exits the application.
func commandExit(commands map[string]cliCommand, cfg *config) error {
	fmt.Println("Exiting the application.")
	os.Exit(0)
	return nil
}
