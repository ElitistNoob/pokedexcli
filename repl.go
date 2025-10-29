package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ElitistNoob/pokedexcli/api"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command := words[0]
		cmd, ok := commands[command]

		if !ok {
			fmt.Println("Unknown Command")
			continue
		}

		if err := cmd.callback(cfg); err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.TrimSpace(strings.ToLower(text))
	list := strings.Split(output, " ")
	return list
}

type config struct {
	apiClient api.Client
	next      *string
	previous  *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations within the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations within the Pokemon world",
			callback:    commandMapBack,
		},
	}
}
