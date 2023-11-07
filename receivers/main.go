package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"
	myVar2 := myStruct{
		FirstName: "Hello",
	}
	log.Println("myVar is set to", myVar.FirstName)
	log.Println("myVar is set by func", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.FirstName)
	log.Println("myVar2 is set by func", myVar2.printFirstName())

}
