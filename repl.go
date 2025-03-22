package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func startRepl() {
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
		command, ok := getCommands()[firstWord]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback()
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
	callback	func() error
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
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}