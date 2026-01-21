package main

import (
	"fmt"
)

// callback for help command
func commandHelp() error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage: \n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil

}
