package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/mtslzr/pokeapi-go"
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

		type pkmnData struct { 
			Pokemon any
			Species any
		}

		var d pkmnData

		p, pErr := pokeapi.Pokemon(pkmnNum)
		if (pErr != nil) {
			log.Fatalln(pErr)
		}

		s, sErr := pokeapi.PokemonSpecies(pkmnNum)
		if (sErr != nil) {
			log.Fatalln(sErr)
		}

		d.Pokemon = p
		d.Species = s

		template.Must(template.ParseFiles("pkmn.html")).Execute(w, d)
	})

	http.ListenAndServe(":8080", nil)
}