package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		words := CleanInput(input)
		cmd, ok := CommandsList[words[0]]
		if ok {
			err := cmd.callback(words[0])
			if err != nil {
				fmt.Println("Unknown command)")
			}
		}
	}
}
