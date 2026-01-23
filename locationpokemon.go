package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/B00m3r0302/PokeDex/internal/pokecache"
)

type PokemonFound []PokemonInfo
type LocationPokemon struct {
	argument string
	cache    *pokecache.Cache
}

type Arguments struct {
	argument []string
}

type PokemonResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon PokemonInfo `json:"pokemon"`
}

type PokemonInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func NewLocationPokemon() *LocationPokemon {
	return &LocationPokemon{
		cache: pokecache.NewCache(10 * time.Second),
	}
}

func LocationAreaShowPokemon(_ string, arguments *Arguments) error {
	if len(arguments.argument) == 0 {
		return fmt.Errorf("please provide a location area name")
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", arguments.argument[0])
	var pokemon PokemonResponse
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range pokemon.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil

}
