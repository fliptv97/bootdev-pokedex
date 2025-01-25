package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fliptv97/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient        pokeapi.Client
	prevLocationAreasUrl *string
	nextLocationAreasUrl *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		parts := cleanInput(scanner.Text())

		if len(parts) == 0 {
			continue
		}

		commandName := parts[0]
		command, ok := commands[commandName]

		if !ok {
			fmt.Println("Unknown command")

			continue
		}

		err := command.callback(cfg)

		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "",
			callback:    commandMapb,
		},
	}
}
