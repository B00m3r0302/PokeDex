package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func resetMapState() {
	mapOffset = 0
	mapResults = nil
	mapTotal = 0
}

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
	resetMapState()
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(_ string) error {
	resetMapState()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, command := range CommandsList {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func GetLocationArea(offset int) (LocationAreaResponse, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?offset=%d&limit=20", offset)

	resp, err := http.Get(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	var apiResp LocationAreaResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return LocationAreaResponse{}, err
	}

	return apiResp, nil
}

func CommandMap(_ string) error {
	// If mapResults is empty, fetch the first page
	if mapOffset == 0 || len(mapResults) == 0 {
		page, err := GetLocationArea(0)
		if err != nil {
			return err
		}
		mapResults = page.Results
		mapTotal = page.Count
	}

	// Calculate how many times to return
	start := mapOffset
	end := mapOffset + 20
	if end > len(mapResults) {
		end = len(mapResults)
	}

	// Print current batch
	for _, loc := range mapResults[start:end] {
		fmt.Println(loc.Name)
	}

	mapOffset = end

	// Fetch next page if needed
	if mapOffset >= len(mapResults) && mapOffset < mapTotal {
		page, err := GetLocationArea(len(mapResults))
		if err != nil {
			return err
		}
		mapResults = append(mapResults, page.Results...)
	}

	// Reset if we've reached the total
	if mapOffset >= mapTotal {
		fmt.Println("Reached all of the location areas.")
		mapOffset = 0
		mapResults = nil
		mapTotal = 0
	}

	return nil
}

func CommandMapB(_ string) error {
	// If mapResults is empty, fetch the first page
	if mapOffset == 0 || len(mapResults) == 0 {
		page, err := GetLocationArea(0)
		if err != nil {
			return err
		}
		mapResults = page.Results
		mapTotal = page.Count
	}

	// Calculate how many times to return
	start := mapOffset - 20
	end := mapOffset + 20
	if end > len(mapResults) {
		end = len(mapResults)
	}

	// Print current batch
	for _, loc := range mapResults[start:end] {
		fmt.Println(loc.Name)
	}

	mapOffset = end

	// Fetch next page if needed
	if mapOffset >= len(mapResults) && mapOffset < mapTotal {
		page, err := GetLocationArea(len(mapResults))
		if err != nil {
			return err
		}
		mapResults = append(mapResults, page.Results...)
	}

	// Reset if we've reached the total
	if mapOffset >= mapTotal {
		fmt.Println("Reached all of the location areas.")
		mapOffset = 0
		mapResults = nil
		mapTotal = 0
	}

	return nil
}

type LocationAreaResponse struct {
	Count   int       `json:"count"`
	Results []NameURL `json:"results"`
}

type NameURL struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var (
	CommandsList map[string]CliCommands

	// State for the "map" command
	mapOffset  int
	mapResults []NameURL
	mapTotal   int
)

type CliCommands struct {
	name        string
	description string
	callback    func(string) error
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
	}
}
