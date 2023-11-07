package main

import (
	"log"

	"github.com/abidahmadq/myprog/helpers"
)

func main() {
	log.Println("Hello")
	var myVar helpers.SomeType
	myVar.TypeName = "Abid"
	log.Println(myVar.TypeName)
}
