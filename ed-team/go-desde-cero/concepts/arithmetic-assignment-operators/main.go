package main

import "fmt"

func main() {
	// arithmetic operators
	var a = (2 + 3) * 5
	fmt.Println(a)
	// Assignment operators
	var b int = 5
	b += b
	fmt.Println(b)
	// Post increment and post decrement (are declaration)
	b++
	a--
	fmt.Println(a, b)
}
