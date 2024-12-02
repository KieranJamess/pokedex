package main

import (
	"bufio"
	"fmt"
	"github.com/KieranJamess/pokedex/internal/pokeapi"
	"os"
	"strings"
	"time"
)

// cliCommand represents a CLI command with its name, description, and a callback function.
type cliCommand struct {
	name        string
	description string
	callback    func(commands map[string]cliCommand, cfg *config, args ...string) error
}

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}

	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the application",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get 20 locations. Use again to display the next further 20 locations.",
			callback:    commandMap,
		},
		"mapback": {
			name:        "mapback",
			description: "Get the previous 20 locations. You can keep using this unless you're on the first page.",
			callback:    commandMapback,
		},
		"explore": {
			name:        "explore",
			description: "Explore a specific area for pokemon found in it. Type 'map' to get a list of areas to explore.",
			callback:    commandExplore,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the CLI tool! Type 'help' to see available commands.")

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := strings.TrimSpace(scanner.Text())
			parts := strings.Fields(input)
			commandName, args := parts[0], parts[1:]
			if cmd, exists := commands[commandName]; exists {
				if err := cmd.callback(commands, cfg, args...); err != nil {
					fmt.Printf("Error executing command: %v\n", err)
				}
			} else {
				fmt.Println("Unknown command. Type 'help' for a list of commands.")
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			break
		}
	}
}
