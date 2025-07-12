package main

import (
	"log"
	"os/exec"
	"strconv"
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

func getLineNumber(in string, searchString string) int64 {
	lineFind := exec.Command("grep", "-n", in, searchString)
	lineFindCmd, err := lineFind.Output()
	if err != nil {
		log.Fatalln("grep line finding error:", err)
	}

	cutLine := exec.Command("cut", "-d", ":", "-f1", string(lineFindCmd))
	cutLineCmd, err := cutLine.Output()
	if err != nil {
		log.Fatalln("Line cut error:", err)
	}

	lineNumber, err := strconv.ParseInt(string(cutLineCmd), 0, 0)
	if err != nil {
		log.Fatalln("Line number conversion error:", err)
	}

	return lineNumber
}
