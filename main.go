package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/thachnguyensg/boot-pokedexcli/internal/pokeapi"
)

var locationEndpoint *pokeapi.LocationEndpoint

func main() {
	locationEndpoint = pokeapi.NewLocationEndpoint("https://pokeapi.co/api/v2/location-area/")
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
				fmt.Println(err.Error())
			}
		}
	}
}
