package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

var CommandsList map[string]CliCommands

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

func CommandExit(_ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(_ string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, command := range CommandsList {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func CommandMap(_ string) error {

}

type CliCommands struct {
	name        string
	description string
	callback    func(string) error
}

type MapLocations struct {
}

func init() {
	CommandsList = map[string]CliCommands{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CommandExit,
		},
		"help": {
			name:        "help",
			description: "Show this help message",
			callback:    CommandHelp,
		},
	}
}
