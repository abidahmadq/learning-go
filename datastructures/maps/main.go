package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	var myString string
	var myInt int

	myString = "Hi"
	myInt = 11

	mySecondString := "another string"
	log.Println(myString, mySecondString, myInt)

	myMap := make(map[string]string)

	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"

	log.Println(myMap["dog"])
	myMap["dog"] = "Fido"
	log.Println(myMap["other-dog"])
	log.Println(myMap["dog"])

	myMap2 := make(map[string]int)
	myMap2["first"] = 1
	myMap2["second"] = 2
	log.Println(myMap2["first"])
	log.Println(myMap2["second"])

	myMap3 := make(map[string]User)
	myMap3["me"] = User{
		FirstName: "Abid",
		LastName:  "Ahmad",
	}
	log.Println(myMap3["me"].FirstName)
	log.Println(myMap3["me"].LastName)

}
