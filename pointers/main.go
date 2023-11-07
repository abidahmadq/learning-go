package main

import "log"

func main() {
	var myString string
	myString = "Green"
	log.Println("mystring is set to", myString)
	changeUsingPointers(&myString)
	log.Println("mystring is set to", myString)

}

func changeUsingPointers(s *string) {
	newValue := "Red"
	log.Println("s is set to", s)
	log.Println("newValue is set to", &newValue)
	*s = newValue
	log.Println("s is now set to", s)
}
