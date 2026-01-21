package main

import (
	"fmt"
	"os"
)

// callback for exit command
func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
