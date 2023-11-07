package main

import (
	"log"

	"github.com/abidahmadq/myprog1/helpers"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumer := helpers.RandomNumber(10)
	intChan <- randomNumer
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
