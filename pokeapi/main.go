package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/mtslzr/pokeapi-go"
)

func capitalize(s string) string {
	return fmt.Sprintf("%v%v", strings.ToUpper(string(s[0])), s[1:])
}

func main() {
	dex, dexErr := pokeapi.Pokedex("national")
	if (dexErr != nil) {
		log.Fatalln(dexErr)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template.Must(template.ParseFiles("main.html")).Execute(w, dex.PokemonEntries)
	})

	http.ListenAndServe(":8080", nil)
}