package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/mtslzr/pokeapi-go"
)

func parseTemp(f string) *template.Template {
	return template.Must(template.ParseFiles(f))
}

func main() {
	dex, dexErr := pokeapi.Pokedex("national")
	if (dexErr != nil) {
		log.Fatalln(dexErr)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template.Must(template.ParseFiles("main.html")).Execute(w, dex)
	})

	http.HandleFunc("/pkmn/{num}", func(w http.ResponseWriter, r *http.Request) {
		pkmnNum := r.PathValue("num")

		type PkmnData struct {
			Pokemon any
			PokemonSpecies any
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

		parseTemp("pkmn.html").Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}