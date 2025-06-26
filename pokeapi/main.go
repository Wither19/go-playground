package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func main() {
	dex, dexErr := pokeapi.Pokedex("national")
	if (dexErr != nil) {
		log.Fatalln(dexErr)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template.Must(template.ParseFiles("main.html")).Execute(w, dex)
	})

	http.HandleFunc("/pkmn/{num}", func(w http.ResponseWriter, r *http.Request) {
		pkmnNum := r.PathValue("num")

		type PkmnData struct {
			Pokemon structs.Pokemon
			PokemonSpecies structs.PokemonSpecies
		}

		p, pErr := pokeapi.Pokemon(pkmnNum)
		if (pErr != nil) {
			log.Fatalln(pErr)
		}

		s, sErr := pokeapi.PokemonSpecies(pkmnNum)
		if (sErr != nil) {
			log.Fatalln(sErr)
		}

		data := PkmnData{Pokemon: p, PokemonSpecies: s}

		template.Must(template.ParseFiles("pkmn.html")).Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}