package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			fmt.Println()
			return
		}
		input := scanner.Text()
		words := CleanInput(input)
		if len(words) == 0 {
			continue
		}

		cmd, ok := CommandsList[words[0]]
		if ok {
			if err := cmd.callback(words[0]); err != nil {
				fmt.Println("Error:", err)
			}
			// do NOT print anything here if err is nil
		} else {
			fmt.Println("Unknown command:", words[0]) // only print if command is not found
		}
	}
}
