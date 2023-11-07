package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	var s2 = "six"
	s := "eight"
	log.Println("s is ", s)
	log.Println("s2 is ", s2)
	log.Println(saySomething("xxx"))

	user := User{
		FirstName: "Abid",
		LastName:  "Ahmad",
	}
	log.Println(user.FirstName)
	log.Println(user.LastName)
	log.Println(user)

}

func saySomething(s string) (string, string) {
	log.Println("s from the saySomething func is", s)
	return s, "World"
}
