package main

import "fmt"

func main() {
	var numbers = [...]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	numberArrSlice := numbers[0:3]
	// modification
	numberArrSlice[1] = "2000000"
	// from the beginning is not necessary to specify the initial index
	numberArrSampleSlice := numbers[:5]
	// same if you want all
	numberArrAllSlice := numbers[:]
	fmt.Println(numbers, numberArrSlice, numberArrSampleSlice, numberArrAllSlice)
	// slice length
	fmt.Println(len(numberArrSampleSlice))
	// capacity
	fmt.Println(cap(numberArrSampleSlice))
	numberArrSlice = append(numberArrSlice, "9", "8", "7", "6", "5", "4", "3", "2", "1", "0", "-1")
	fmt.Println(numbers, numberArrSlice, numberArrSampleSlice, numberArrAllSlice)
	// Create new arrays from 0
	pets := make([]string, 2, 3)
	pets = append(pets, "Dog", "Cat", "Bird")
	fmt.Println(pets)
	// nil value
	var dogs []string
	fmt.Println("nill value? ", dogs == nil)
}
