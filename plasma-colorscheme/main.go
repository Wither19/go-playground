package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatalln("Window creation error", err)
	}

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.SetDefaultSize(800, 600)

	mainGrid, err := gtk.GridNew()
	if err != nil {
		log.Fatalln("Grid creation error", err)
	}

	colorSchemeSelect, err := gtk.ComboBoxTextNew()
	if err != nil {
		log.Fatalln("Combo Box Text creation error:", err)
	}

	for i, colorScheme := range getColorschemeList() {
		colorSchemeSelect.Append(fmt.Sprintf("colorscheme-%d", i), colorScheme.Name)
	}

	mainGrid.Attach(colorSchemeSelect, 1, 1, 1, 1)

	win.Add(mainGrid)
	win.ShowAll()

	gtk.Main()
}