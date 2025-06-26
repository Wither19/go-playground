package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
)

func parseTemplate(filename string) *template.Template {
	tmpl := template.Must(template.ParseFiles(filename))
	
	return tmpl
}

func main() {

type Option struct {
	Text string `json:"text"`
	Chapter  string `json:"arc"`
}

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Story map[string]Chapter

	fileName := flag.String("file", "./gopher.json", "The JSON file for the story")
	flag.Parse()

	storyJson, jsonErr := os.Open(*fileName)
	if (jsonErr != nil) {
		log.Fatal(jsonErr)
	}

	decodedStory := json.NewDecoder(storyJson)
	var story Story
	if err := decodedStory.Decode(&story); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		parseTemplate("temp.html").Execute(w, story["intro"])
	})
	http.ListenAndServe(":8080", nil)
}