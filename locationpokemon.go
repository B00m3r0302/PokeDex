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
	Pokemon []PokemonInfo `json:"pokemon"`
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
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", "1")
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

	fmt.Println(PokemonInfo{})
	return nil

}
