package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mrmaths1212/pokedexcli/internal/pokeapi"
)

func cleanInput(input string) []string {
	lower := strings.ToLower(input)
	words := strings.Fields(lower)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	pokeApiClient pokeapi.Client
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show available commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show the command map",
			callback:    commandMap,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			continue
		}
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			fmt.Println("No valid input provided.")
			continue
		}
		command, exist := getCommand()[cleaned[0]]
		if !exist {
			fmt.Printf("Unknown command: %s\n", cleaned[0])
			continue
		}
		err := command.callback()
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			continue
		}

	}
}
