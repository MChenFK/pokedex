package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"

	"github.com/MChenFK/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient		pokeapi.Client
	nextLocationsURL	*string
	prevLocationsURL	*string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}

		firstWord := cleanedInput[0]
		args := []string{}
		if len(cleanedInput) > 1 {
			args = cleanedInput[1:]
		}

		command, ok := getCommands()[firstWord]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		}

		
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

type cliCommand struct {
	name		string
	description	string
	callback	func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Takes the name of a location area and displays a list of all the Pokemon located there",
			callback:    commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}