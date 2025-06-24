package main

import (
	"fmt"
	"log"

	"github.com/mtslzr/pokeapi-go"
)

func main() {
	natlDex, dexFetchErr := pokeapi.Pokedex("national")
	if (dexFetchErr != nil) {
		log.Fatal(dexFetchErr)
	}

	fmt.Println(natlDex.PokemonEntries[250].PokemonSpecies.Name)
}