package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/JoshGuarino/PokeGo/pkg/resources/games"
	"github.com/JoshGuarino/PokeGo/pkg/resources/pokemon"
)

func main() {
	game := games.NewGamesGroup()
	
	dex, dexErr := game.GetPokedex("national")
	if (dexErr != nil) {
		log.Fatalln(dexErr)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template.Must(template.ParseFiles("main.html")).Execute(w, dex)
	})

	http.HandleFunc("/pkmn/{num}", func(w http.ResponseWriter, r *http.Request) {
		pkmnNum := r.PathValue("num")

		type PkmnData struct {
			Pokemon any
			PokemonSpecies any
		}
		
		pkmn := pokemon.NewPokemonGroup()

		p, pErr := pkmn.GetPokemon(pkmnNum)
		if (pErr != nil) {
			log.Fatalln(pErr)
		}

		s, sErr := pkmn.GetPokemonSpecies(pkmnNum)
		if (sErr != nil) {
			log.Fatalln(sErr)
		}

		data := PkmnData{Pokemon: p, PokemonSpecies: s}

		template.Must(template.ParseFiles("pkmn.html")).Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}