package main

import (
	"fmt"
	"os"
	"strings"
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

func CommandHelp(text string) error {
	if text == "help" {
		fmt.Println(
			"Welcome to the Pokedex!\n" +
				"Usage:\n\n" +
				"help: Displays a help message\n" +
				"exit: Exit the Pokedex")
	}
	return nil
}

type CliCommands struct {
	name        string
	description string
	callback    func(string) error
}

var CommandsList = map[string]CliCommands{
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
