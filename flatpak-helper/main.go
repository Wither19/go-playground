package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	window, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatalln("Window creation error:", err)
	}
	window.SetTitle("Flatpak Helper")
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	optionsContainer, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 2)
	if err != nil {
		log.Fatalln("Box creation error:", err)
	}

	installBtn, err := gtk.ButtonNewWithLabel("Install a Package")
	if err != nil {
		log.Fatalln("Button creation error", err)
	}

	removeBtn, err := gtk.ButtonNewWithLabel("Remove a Package")
	if err != nil {
		log.Fatalln("Button creation error", err)
	}

	optionsContainer.Add(installBtn)
	optionsContainer.Add(removeBtn)

	window.Add(optionsContainer)


	gtk.Main()
}