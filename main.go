package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanned := scanner.Scan()

		if !scanned {
			log.Fatal("Unable to read input")
		}

		line := scanner.Text()
		fmt.Println("input:", line)
		if line == "exit" {
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
	}
}
