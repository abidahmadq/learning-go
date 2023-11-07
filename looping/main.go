package main

import (
	"fmt"
	"log"
)

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i, "th  Hi")
	}

	animals := []string{"dog", "fish", "horse", "cow", "cat"}
	for i, animal := range animals {
		log.Println(i, " is ", animal)
	}

	for _, animal := range animals {
		log.Println(animal)
	}

	animals2 := make(map[string]string)
	animals2["dog"] = "Fido"
	animals2["cat"] = "Fluffy"

	for animalType, animal := range animals2 {
		log.Println(animalType, animal)
	}

	var firstLine = "Once upon a midnight dreary"
	for i, l := range firstLine {
		log.Println(i, " : ", l)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}
	var users []User
	users = append(users, User{"Abid", "Ahmad", "abidahmad@gmail.com", 25})
	users = append(users, User{"Abid1", "Ahmad1", "abidahmad1@gmail.com", 23})
	users = append(users, User{"Abid2", "Ahmad2", "abidahmad2@gmail.com", 22})
	users = append(users, User{"Abid2", "Ahmad3", "abidahmad3@gmail.com", 21})
	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Age, l.Email)
	}

	dog := Dog{
		Name:  "Samson",
		Breed: "German Shephered",
	}
	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "Grey",
		NumberOfTeeth: 38,
	}
	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "Ugh"
}

func (d *Gorilla) NumberOfLegs() int {
	return 4
}
