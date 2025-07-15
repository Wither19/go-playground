package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {

type Chapter struct {
	Title   	string   `json:"title"`
	Reference string `json:"reference"`
	Story   []string `json:"story"`
	Options []struct {
		Text 		 	string `json:"text"`
		Chapter  	string `json:"arc"`
	} `json:"options"`
}

	fileName := flag.String("file", "./gopher.json", "The JSON file for the story")
	flag.Parse()

	storyJson, jsonErr := os.Open(*fileName)
	if (jsonErr != nil) {
		log.Fatal(jsonErr)
	}

	var story map[string]Chapter

	decodedStory := json.NewDecoder(storyJson)
	if err := decodedStory.Decode(&story); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/intro", http.StatusMovedPermanently)
	})

	for _, chapter := range story {

		http.HandleFunc(fmt.Sprintf("/%v", chapter.Reference), func(w http.ResponseWriter, r *http.Request) {
			template.Must(template.ParseFiles("temp.html")).Execute(w, chapter)
		})
	}

	http.ListenAndServe(":8080", nil)
}