package main

import "fmt"

func main() {
	var color string = "cyan"
	// Address of Operator
	fmt.Printf("Type: %T, Value: %s, Allocaiton Memory Direction: %v\n", color, color, &color)
	// Pointer
	var pointer *string
	pointer = &color
	fmt.Printf("Type: %T, Value: %v\n", pointer, pointer)
	// The pointer prints the same memory allocation of color
	// To get the value which a pointer is referenced to:
	var value string = *pointer
	fmt.Printf("Type: %T, Value: %v\n", value, value)
	// updating the original value from the pointer
	*pointer = "green"
	fmt.Printf("Type: %T, Value: %v\n", color, color)
}
