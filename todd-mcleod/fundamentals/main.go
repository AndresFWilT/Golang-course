package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println("a is", a)

	// printing with blank value
	a, b, c, _, e := 1, 2, 3, 4, "Happiness"
	fmt.Println(a, b, c, e)

	// zero values
	var zeroInt int
	var zeroString string

	fmt.Println(zeroInt, zeroString)
	// print 0, ""

	//using printf to binary and hexadecimal
	fmt.Printf("%v \t %b \t %#X\n", a, a, a)
	fmt.Printf("%v \t %b \t %#X\n", b, b, b)
	fmt.Printf("%v \t %b \t %#X\n", c, c, c)
}
