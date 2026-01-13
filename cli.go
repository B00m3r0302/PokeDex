package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/B00m3r0302/PokeDex/internal/pokecache"
)

const pageSize = 20

type MapNavigator struct {
	offset int
	total  int
	cache  *pokecache.Cache
}

func NewMapNavigator() *MapNavigator {
	return &MapNavigator{
		cache: pokecache.NewCache(10 * time.Second),
	}
}

func (m *MapNavigator) Reset() {
	m.offset = 0
	m.total = 0
}

func (m *MapNavigator) MoveForward() error {
	page, err := GetLocationArea(m.offset)
	if err != nil {
		return err
	}

	if m.offset == 0 {
		m.total = page.Count
	}

	for _, loc := range page.Results {
		fmt.Println(loc.Name)
		m.cache.Add(loc.Name, []byte(loc.Name))
	}

	m.offset += pageSize

	if m.offset >= m.total {
		fmt.Println("Reached the end of the location areas.")
		m.Reset()
	}

	return nil
}

func (m *MapNavigator) MoveBackward() error {
	if m.offset < pageSize {
		fmt.Println("Already at the beginning of the location areas.")
		return nil
	}

	m.offset -= pageSize

	page, err := GetLocationArea(m.offset)
	if err != nil {
		return err
	}

	if m.total == 0 {
		m.total = page.Count
	}

	for _, loc := range page.Results {
		fmt.Println(loc.Name)
		m.cache.Add(loc.Name, []byte(loc.Name))
	}

	return nil
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
	navigator.Reset()
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(_ string) error {
	navigator.Reset()
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
	return navigator.MoveForward()
}

func CommandMapB(_ string) error {
	// If mapResults is empty, fetch the first page
	return navigator.MoveBackward()
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
	navigator    *MapNavigator
)

type CliCommands struct {
	name        string
	description string
	callback    func(string) error
}

func init() {

	navigator = NewMapNavigator()

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
