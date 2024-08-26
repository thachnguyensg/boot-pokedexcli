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
