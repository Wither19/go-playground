package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/mtslzr/pokeapi-go/structs"
)

func parseTemp(f string) *template.Template {
	return template.Must(template.ParseFiles(f))
}

func getAPILink(cat string, id string) string {
	return fmt.Sprintf("api-data/%v/%v/index.json", cat, id)
}

func leadingZeroes(num int, length int) string {
	return fmt.Sprintf("%0*d", length, num)
}

func getNatlDex() structs.Pokedex {
	dexURL := getAPILink("pokedex", "1")

	dex, err := os.ReadFile(dexURL)

	if (err != nil) {
		log.Fatalln("Dex fetch error:", err)
	}

	var pokedex structs.Pokedex

	dexUnpackErr := json.Unmarshal(dex, &pokedex)

	if (dexUnpackErr != nil) {
		log.Fatalln("Dex unpack error:", err)
	}

	return pokedex
}

func getPkmn(id string) structs.Pokemon {
	pURL := getAPILink("pokemon", id)

	p, pErr := os.ReadFile(pURL)

	if (pErr != nil) {
		log.Fatalln("Pokemon fetch error:", pErr)
	}

	var pkmn structs.Pokemon

	pUnpackErr := json.Unmarshal(p, &pkmn)

	if (pUnpackErr != nil) {
		log.Fatalln("Pokemon unpack error:", pUnpackErr)
	}
	return pkmn
}

func getPkmnSpecies(id string) structs.PokemonSpecies {
	sURL := getAPILink("pokemon-species", id)
	
	s, sErr := os.ReadFile(sURL)

	if (sErr != nil) {
		log.Fatalln("Species fetch error:", sErr)
	}

	var species structs.PokemonSpecies

	sUnpackErr := json.Unmarshal(s, &species)

	if (sUnpackErr != nil) {
		log.Fatalln("Species unpack error:", sUnpackErr)
	}	
	return species
}

func mainPageHandle(w http.ResponseWriter, r *http.Request) {
	d := getNatlDex().PokemonEntries

	parseTemp("main.html").Execute(w, d)
}

func pkmnLoadfunc(w http.ResponseWriter, r *http.Request) {
	pkmnID := r.PathValue("id")

	type PkmnData struct {
		Pokemon structs.Pokemon
		PokemonSpecies structs.PokemonSpecies
		paddedID string
	}

	pkmn := getPkmn(pkmnID)
	species := getPkmnSpecies(pkmnID)

	data := PkmnData{Pokemon: pkmn, PokemonSpecies: species, paddedID: leadingZeroes(pkmn.ID, 4)}

	parseTemp("pkmn.html").Execute(w, data)

}

func main() {
	sassSource := "./static/scss/App.scss"
	newCss := "./static/css/style.css"
	sassBuild := exec.Command("sass", sassSource, newCss, "--no-source-map")

	if err := sassBuild.Run(); err != nil {
		log.Fatalln("Sass build error:", err)
	} else {
		fmt.Println("Sass successfully transpiled")
	}
	
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", mainPageHandle)
	http.HandleFunc("/pkmn/{id}", pkmnLoadfunc)

	fmt.Println("Server active")

	http.ListenAndServe(":8080", nil)
}

