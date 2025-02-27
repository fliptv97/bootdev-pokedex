package main

import "fmt"

func commandHelp(_ *config, _ ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range getCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}

	return nil
}
