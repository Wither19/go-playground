package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/samber/lo"
)

func getColorschemeList() []ColorScheme {
	colorschemeLS := exec.Command("plasma-apply-colorscheme", "-l")
	execColorSchemeLS, err := colorschemeLS.Output()
	if err != nil {
		log.Fatalln("Color Scheme list error:", err)
	}

	colorSchemeSlice := strings.Split(strings.ReplaceAll(string(execColorSchemeLS), "*", ""), "\n")[1:]
	colorSchemeSlice = colorSchemeSlice[:len(colorSchemeSlice)-1]

	colorSchemeObj := lo.Map(colorSchemeSlice, func(item string, index int) ColorScheme {
		var obj ColorScheme

		obj.Name = item

		if strings.Contains(item, "(current color scheme)") {
			obj.Name = strings.ReplaceAll(item, "(current color scheme)", "")
			obj.Current = true
		}

		obj.Name = strings.TrimSpace(obj.Name)

		return obj
	})

	return colorSchemeObj
}