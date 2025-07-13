package main

import (
	"log"
	"os/exec"
	"strings"
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

