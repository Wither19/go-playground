package main

import (
	"fmt"
	"log"
)

func pkgInstall() {
	log.Fatalln("That feature has not been implemented yet")
}

func pkgRemove() {
	fmt.Print(sliceFlatpakList(flatpakList("name")))
}