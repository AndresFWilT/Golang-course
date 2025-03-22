package main

import "fmt"

func main() {
	// defer works as a stack
	defer fmt.Println("3")
	defer fmt.Println("2")
	defer fmt.Println("1")

	// other sample
	var number uint8 = 2
	defer fmt.Printf("Value: %v\t Type: %T\n", number, number)

	number++
	fmt.Printf("Value: %v\t Type: %T\n", number, number)
}
