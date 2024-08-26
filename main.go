package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	initCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanned := scanner.Scan()

		if !scanned {
			log.Fatal("Unable to read input")
		}

		line := scanner.Text()

		if command, ok := commands[line]; ok {
			err := command.callback()
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}
}
