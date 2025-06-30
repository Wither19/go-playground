package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/mtslzr/pokeapi-go/structs"
)

func parseTemp(f string) *template.Template {
	return template.Must(template.ParseFiles(f))
}

func getAPILink(cat string, id string) string {
	return fmt.Sprintf("api-data/data/api/v2/%v/%v/index.json", cat, id)
}

func mainPageHandle(w http.ResponseWriter, r *http.Request) {
	dexURL := getAPILink("pokedex", "1")

	dex, err := os.ReadFile(dexURL)

	if (err != nil) {
		log.Fatalln("Dex error:", err)
	}

	var pokedex structs.Pokedex

	dexUnpackErr := json.Unmarshal(dex, &pokedex)

	if (dexUnpackErr != nil) {
		log.Fatalln("Dex unpacking error:", err)
	}

	parseTemp("main.html").Execute(w, pokedex.PokemonEntries)
}

func pkmnLoadfunc(w http.ResponseWriter, r *http.Request) {
	pkmnNum := r.PathValue("num")

	type PkmnData struct {
		Pokemon structs.Pokemon
		PokemonSpecies structs.PokemonSpecies
	}
	pURL := getAPILink("pokemon", pkmnNum)

	p, pErr := os.ReadFile(pURL)

	if (pErr != nil) {
		log.Fatalln("Pokemon error:", pErr)
	}

	sURL := getAPILink("pokemon-species", pkmnNum)
	
	s, sErr := os.ReadFile(sURL)

	if (sErr != nil) {
		log.Fatalln("Species error:", sErr)
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

	http.ListenAndServe(":8080", nil)
}

