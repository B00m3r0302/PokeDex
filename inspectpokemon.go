package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type BasePokemonInfo struct {
	Name   string  `json:"name"`
	Height int     `json:"height"`
	Weight int     `json:"weight"`
	Stats  []Stats `json:"stats"`
	Types  []Types `json:"types"`
}

type Stats struct {
	Name  StatsName `json:"stat"`
	Value int       `json:"base_stat"`
}

type StatsName struct {
	Name string `json:"name"`
}

type Types struct {
	Type TypeName `json:"type"`
}

type TypeName struct {
	Name string `json:"name"`
}

func (p *PokeDex) GetPokemon(name string) error {
	_, value := p.cache.Get(name)
	if value == false {
		return fmt.Errorf("Pokemon not found in your pokedex")
	}

	return nil
}

func InspectPokemon(_ string, arguments *Arguments) error {
	if len(arguments.argument) == 0 || arguments.argument[0] == "" {
		return fmt.Errorf("Please enter a pokemon name to inspect it...")
	}

	check_pokedex := pokeDex.GetPokemon(arguments.argument[0])
	if check_pokedex != nil {
		return check_pokedex
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", arguments.argument[0])
	var pokemon BasePokemonInfo
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Pokemon not found \n%v", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("     -%s: %d\n", stat.Name.Name, stat.Value)
	}
	fmt.Println("Types:")
	for _, types := range pokemon.Types {
		fmt.Printf("     -%s\n", types.Type.Name)
	}

	return nil
}
