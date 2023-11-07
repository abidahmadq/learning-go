package main

import "fmt"

func main() {
	fmt.Println("Hello, World!!")
	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel World"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to", i)
	whatWasSaid, somethingElseSaid := saySomething()
	fmt.Println(whatWasSaid, somethingElseSaid)
}

func saySomething() (string, string) {
	return "something", "notsomething"
}
