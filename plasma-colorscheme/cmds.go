package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/samber/lo"
)

func getPlasmaColorSchemes(r bool) []string {
	colorschemeLS := exec.Command("plasma-apply-colorscheme", "-l")

	execColorSchemeLS, err := colorschemeLS.Output()
	if err != nil {
		log.Fatalln("Color Scheme list error:", err)
	}

	stringList := strings.Split(string(execColorSchemeLS), "\n")

	stringList = stringList[1:len(stringList)-1]


	if r {
		stringList = lo.Map(stringList, func(item string, i int) string {
			return strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(item, "(current color scheme)", ""), "*", ""))
		})
	} else {
			stringList = lo.Map(stringList, func(item string, i int) string {
				return strings.TrimSpace(strings.ReplaceAll(item, "*", ""))
			})
	}

	return stringList
}

func getActiveColorScheme() string {
	colorSchemes := getPlasmaColorSchemes(false)
	activeColorScheme := ""

	for _, colorScheme := range colorSchemes {
		if strings.Contains(colorScheme, "(current color scheme)") {
			activeColorScheme = strings.TrimSpace(strings.ReplaceAll(colorScheme, "(current color scheme)", ""))
		}
	}

	return activeColorScheme

}

func plasmaColorSchemeChange(colorScheme string) {
	changeCmd := exec.Command("plasma-apply-colorscheme", colorScheme)

		if err := changeCmd.Run(); err != nil {
			log.Fatalln("Plasma color scheme change error:", err)
		}
}

func breezeModeToggle() {
	theme := "BreezeLight"
	if a := getActiveColorScheme(); strings.Contains(strings.ToLower(a), "light") {
		theme = "BreezeDark"
	}

	changeCmd := exec.Command("plasma-apply-colorscheme", theme)

	if err := changeCmd.Run(); err != nil {
		log.Fatalln("Plasma color scheme change error:", err)
	}

}