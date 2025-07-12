package main

import (
	"fmt"
	"log"
	"strconv"
)


func main() {

	mainOptions := []string{
		"Install a package",
		"Remove a package",
	}

	var userChoice string

	for i, option := range mainOptions {

		fmt.Printf("%d. %v\n", i + 1, option)
	}

	fmt.Print("\nWhat would you like to do? [1-4] ")
	fmt.Scanln(&userChoice)

	choiceNumber, err := strconv.ParseInt(userChoice, 0, 0)
	if err != nil {
		log.Fatalln("Selected option is not a number")
	}

	switch choiceNumber {
	case 1:
		pkgInstall()
	case 2:
		pkgRemove()
	} 

}