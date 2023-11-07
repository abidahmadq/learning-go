package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice []string
	mySlice = append(mySlice, "Abid", "Ahmad")
	mySlice = append(mySlice, "Abid", "Ahmad")
	mySlice = append(mySlice, "Abid", "Ahmad")
	mySlice = append(mySlice, "Abid", "Ahmad")

	log.Println(mySlice)
	var mySlice1 []int
	mySlice1 = append(mySlice1, 1)
	mySlice1 = append(mySlice1, 2)
	mySlice1 = append(mySlice1, 3)
	mySlice1 = append(mySlice1, 4)
	log.Println(mySlice1)
	sort.Slice(mySlice1, func(i, j int) bool {
		return i > j
	})
	log.Println(mySlice1)
	sort.Ints(mySlice1)
	log.Println(mySlice1)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Println(numbers)
	log.Println(numbers[5])
	log.Println(numbers[6:9])
	log.Println(numbers[:])

	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	log.Println(numbers)

	names := []string{"one", "two", "fish", "cat"}
	log.Println(names)

}
