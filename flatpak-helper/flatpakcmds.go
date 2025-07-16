package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/gotk3/gotk3/gtk"
)

func flatpakList(column string) string {
	var flatpakListCmd *exec.Cmd
	if (column != "") {
		flatpakListCmd = exec.Command("flatpak", "list", "--columns", column)
	} else {
		flatpakListCmd = exec.Command("flatpak", "list")
	}

	list, err := flatpakListCmd.Output()
	if err != nil {
		log.Fatalln("Flatpak listing error:", err)
	}

	return string(list)
} 

func sliceFlatpakList(list string) []string {
	return strings.Split(list, "\n")
} 


func pkgRemove(pkgID string) {
	flatpakRemoveCmd := exec.Command("flatpak", "remove", pkgID, "-y")
	flatpakRemoveExec, err := flatpakRemoveCmd.CombinedOutput()

	removalModal, _ := gtk.DialogNew()

	if err != nil {
		removalModal.SetTitle("Package Removal Failed")
	} else {
		removalModal.SetTitle(string(flatpakRemoveExec))
	}
		
		removalModal.AddButton("OK", gtk.RESPONSE_OK)

		removalModal.Connect("response", func() {
			gtk.MainQuit()
		})
}