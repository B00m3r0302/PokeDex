package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/B00m3r0302/PokeDex/internal/pokecache"
)

type PokeDex struct {
	argument string
	cache    *pokecache.Cache
}

func NewPokeDex() *PokeDex {
	return &PokeDex{
		cache: pokecache.NewCache(10 * time.Second),
	}
}

type CatchCommands struct {
	argument []string
}

type PokemonDetails struct {
	BaseExperience int `json:"base_experience"`
}

type Pokemon struct {
	name string
}

func (p *PokeDex) AddPokemon(argument string) {
	p.cache.Add(argument, []byte(argument))
}

func CatchPokemon(_ string, arguments *Arguments) error {
	if len(arguments.argument) == 0 {
		return fmt.Errorf("Please enter a pokemon name to try and catch it...")
	}

	throw := fmt.Sprintf("Throwing a Pokeball at %s...", arguments.argument[0])
	fmt.Println(throw)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", arguments.argument[0])
	var response PokemonDetails
	randomNumber := rand.Intn(100)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	if randomNumber < response.BaseExperience {
		fmt.Printf("%s escaped!\n", arguments.argument[0])
		return nil
	}

	pokeDex.AddPokemon(arguments.argument[0])
	fmt.Printf("%s was caught!\n", arguments.argument[0])
	return nil
}
