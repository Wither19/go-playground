package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"golang.org/x/net/html"
)


func main() {
	fileFlag := flag.String("file", "example.html", "The HTML file to read")
	flag.Parse()

	htmlOpen, htmlOpenErr := os.ReadFile(*fileFlag)
	if (htmlOpenErr != nil) {
		log.Fatal(htmlOpenErr)
	}

	parsedDoc, parseErr := html.Parse(strings.NewReader(string(htmlOpen)))
	if (parseErr != nil) {
		log.Fatal(parseErr)
	}

	for t := range parsedDoc.Descendants() {
		if (slices.ContainsFunc(t.Attr, func(a html.Attribute) bool {
			return a.Key == "href"
		})) {
			// Only executes for tags with "href" attributes
			fmt.Printf("\n")
		}
	}
}