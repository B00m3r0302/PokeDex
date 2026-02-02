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

func CommandExit(_ string, arguments *Arguments) error {
	locationavigator.Reset()
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(_ string, arguments *Arguments) error {
	locationavigator.Reset()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, command := range CommandsList {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func CommandExplore(text string, arguments *Arguments) error {
	return nil
}

func CommandMap(_ string, arguments *Arguments) error {
	return locationavigator.LocationAreaMoveForward()
}

func CommandMapB(_ string, arguments *Arguments) error {
	// If mapResults is empty, fetch the first page
	return locationavigator.LocationAreaMoveBackward()
}

var (
	CommandsList     map[string]CliCommands
	locationavigator *MapNavigator
	locationpokemon  *LocationPokemon
	pokeDex          *PokeDex
)

type CliCommands struct {
	name        string
	description string
	callback    func(string, *Arguments) error
}

func init() {

	locationavigator = NewMapNavigator()
	locationpokemon = NewLocationPokemon()
	pokeDex = NewPokeDex()

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
		"map": {
			name:        "map",
			description: "Get location areas",
			callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Move back in location areas",
			callback:    CommandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Get Pokemon in location areas",
			callback:    LocationAreaShowPokemon,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback:    CatchPokemon,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon",
			callback:    InspectPokemon,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Get a list of pokemon in the pokedex",
			callback:    ListPokemon,
		},
	}
}
