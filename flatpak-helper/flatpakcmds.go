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
	if err := flatpakRemoveCmd.Start(); err != nil {
		errorModal, _ := gtk.DialogNew()

		errorModal.SetTitle("Package Removal Failed")
		errorModal.AddButton("OK", gtk.RESPONSE_OK)

		errorModal.Connect("response", func() {
			gtk.MainQuit()
		})
	} else {

	}
	_ = flatpakRemoveCmd.Wait()
}