package main

import (
	"strings"
	"fmt"
	"os"
)

func CleanInput(text string) []string {
	list := strings.Fields(text)
	finalList := []string{}
	for i := range list {
		list[i] = strings.TrimSpace(list[i])
		list[i] = strings.ToLower(list[i])
		finalList = append(finalList, list[i])
	}
	return finalList
}

func CommandExit(text string) error {
	if text == "exit" {
		fmt.Println("Closing the Pokedex... Goodbye!")
		os.Exit(0)
	}
	return nil
}
