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

	mainContainer, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	if err != nil {
		log.Fatalln("Main container creation error:", err)
	}

	windowHeaderBar, err := gtk.HeaderBarNew()
	if err != nil {
		log.Fatalln("Header bar creation error:", err)
	}

	windowHeaderBar.SetTitle("What would you like to do?")
	
	mainContainer.Add(windowHeaderBar)

	optionsContainer, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	if err != nil {
		log.Fatalln("Box creation error:", err)
	}

	installBtn, err := gtk.ButtonNewWithLabel("Install a Package")
	if err != nil {
		log.Fatalln("Button creation error", err)
	}

	// installBtn.Connect("clicked", pkgInstall)

	removeBtn, err := gtk.ButtonNewWithLabel("Remove a Package")
	if err != nil {
		log.Fatalln("Button creation error", err)
	}
	
	removeBtn.Connect("clicked", func() {

		windowHeaderBar.SetTitle("Select a package to remove")

		packagesList := getPkgList(window)

		mainContainer.Add(packagesList)
		packagesList.ShowAll()

		installBtn.Hide()
		removeBtn.Hide()
	})

	installBtn.SetHExpand(true)
	removeBtn.SetHExpand(true)

	installBtn.SetMarginTop(20)

	installBtn.SetMarginStart(40)
	installBtn.SetMarginEnd(40)

	removeBtn.SetMarginStart(40)
	removeBtn.SetMarginEnd(40)

	optionsContainer.Add(installBtn)
	optionsContainer.Add(removeBtn)

	mainContainer.Add(optionsContainer)

	window.Add(mainContainer)

	window.SetDefaultSize(640, 480)

	window.ShowAll()

	gtk.Main()
}