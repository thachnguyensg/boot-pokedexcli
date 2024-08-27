package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	callback func() error
	name     string
	desc     string
}

var commands map[string]cliCommand

func initCommands() {
	commands = map[string]cliCommand{
		"help": {
			callback: help,
			name:     "help",
			desc:     "Display help",
		},
		"exit": {
			callback: exit,
			name:     "exit",
			desc:     "Exit the Pokedex",
		},
		"map": {
			callback: mapCommand,
			name:     "map",
			desc:     "Displays the names of 20 location areas in the Pokemon world",
		},
		"mapb": {
			callback: mapbCommand,
			name:     "mapb",
			desc:     "Displays the names of previous 20 location areas in the Pokemon world",
		},
	}
}

func help() error {
	fmt.Println("Pokedex")
	fmt.Println("")
	fmt.Println("Available commands:")
	fmt.Println("")
	for _, c := range commands {
		fmt.Printf("%s: %s\n", c.name, c.desc)
	}
	fmt.Println("")
	return nil
}

func exit() error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}

func mapCommand() error {
	if locationEndpoint == nil {
		return fmt.Errorf("Location endpoint is not initialized")
	}

	res, err := locationEndpoint.GetNext()
	if err != nil {
		return err
	}

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func mapbCommand() error {
	if locationEndpoint == nil {
		return fmt.Errorf("Location endpoint is not initialized")
	}

	res, err := locationEndpoint.GetPrev()
	if err != nil {
		return err
	}

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	return nil
}
