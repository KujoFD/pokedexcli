package main

import (
	"fmt"
	"os"
)

// callback for exit command
func commandExit(cfg *config, args ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
