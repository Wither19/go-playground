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

	win.SetTitle("Plasma Color Scheme Changer")

	win.SetIconName("org.gnome.ColorViewer")
	win.SetDefaultSize(640, 480)

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
	colorSchemeSelect.SetActive(0)

	colorSchemeSelect.Connect("changed", func() {
		themeChangeBtn.SetSensitive(!(activeTheme == colorSchemeSelect.GetActiveText()))
	})

	themeChangeBtn.Connect("clicked", func() {
		activeTheme = colorSchemeSelect.GetActiveText()
		plasmaColorSchemeChange(activeTheme)
		themeChangeBtn.SetSensitive(false)
	})

	breezeToggleBtn, err := gtk.ButtonNewWithLabel("Toggle Light / Dark mode")
	if err != nil {
		log.Fatalln("Button creation error:", err)
	}

	breezeToggleBtn.Connect("clicked", func() {
		breezeModeToggle()
		activeTheme = getActiveColorScheme()

	})

	mainGrid.Attach(colorSchemeSelect, 0, 0, 1, 1)
	mainGrid.Attach(themeChangeBtn, 2, 0, 2, 1)
	mainGrid.Attach(breezeToggleBtn, 0, 2, 1, 1)

	win.Add(mainGrid)
	win.ShowAll()

	gtk.Main()
}
