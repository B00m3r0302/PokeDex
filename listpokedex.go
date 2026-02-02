package main

import (
	"fmt"
)

func ListPokemon(_ string, arguments *Arguments) error {
	key := pokeDex.cache.Keys()
	if len(key) == 0 {
		return fmt.Errorf("No Pokemon in the Pokedex...")
	}

	fmt.Println("List of Pokemons")
	for _, name := range key {
		fmt.Printf("-%s\n", name)
	}
	return nil
}
