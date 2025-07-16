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

	themeList := getPlasmaColorSchemes(true)
	activeTheme := getActiveColorScheme()

	for i, colorScheme := range themeList {
		colorSchemeSelect.Append(fmt.Sprintf("colorscheme-%d", i), colorScheme)
	}

	themeChangeBtn, err := gtk.ButtonNewWithLabel("Apply")
	if err != nil {
		log.Fatalln("Button creation error:", err)
	}

	colorSchemeSelect.Connect("changed", func() {
		selectedTheme := colorSchemeSelect.GetActiveText()

		themeChangeBtn.SetSensitive(!(activeTheme == selectedTheme))
	})

	mainGrid.Attach(colorSchemeSelect, 1, 1, 1, 1)
	mainGrid.Attach(themeChangeBtn, 6, 6, 2, 1)

	win.Add(mainGrid)
	win.ShowAll()

	gtk.Main()
}