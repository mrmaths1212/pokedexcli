package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, comm := range getCommand() {
		fmt.Printf("%s: %s\n", comm.name, comm.description)
	}
	fmt.Println()
	return nil
}
