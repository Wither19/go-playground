package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func parseTemp(f string) *template.Template {
	return template.Must(template.ParseFiles(f))
}

func mainPageHandle(w http.ResponseWriter, r *http.Request) {
	dex, dexErr := pokeapi.Pokedex("national")
	if (dexErr != nil) {
		log.Fatalln(dexErr)
	}

	parseTemp("main.html").Execute(w, dex.PokemonEntries)
}

func pkmnLoadfunc(w http.ResponseWriter, r *http.Request) {
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

	if (pErr == nil && sErr == nil) {
		data := PkmnData{Pokemon: p, PokemonSpecies: s}

		parseTemp("pkmn.html").Execute(w, data)
	} else {
		http.Redirect(w, r, "/home", http.StatusNotFound)
	}

}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/home", mainPageHandle)
	http.HandleFunc("/pkmn/{num}", pkmnLoadfunc)

	http.ListenAndServe("localhost:8080", nil)
}

