package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/mtslzr/pokeapi-go"
)

func capitalize(s string) string {
	return fmt.Sprintf("%v%v", strings.ToUpper(string(s[0])), s[1:])
}

func main() {
	var dexNum string
	fmt.Print("Get info on a Pokémon by its Pokédex # ")
	fmt.Scanln(&dexNum)

	pkmn, pkmnErr := pokeapi.Pokemon(dexNum)
	if (pkmnErr != nil) {
		log.Fatalln(pkmnErr)
	}

	fmt.Printf("#%d %v\n", pkmn.ID, capitalize(pkmn.Name))
	fmt.Print("Types: ")
}