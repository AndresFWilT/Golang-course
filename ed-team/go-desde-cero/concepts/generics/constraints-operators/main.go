package main

import "fmt"

func main() {
	var result bool
	result = Includes([]string{"a", "b"}, "c")
	fmt.Println(result)
	result = Includes([]string{"a", "b", "c", "d"}, "d")
	fmt.Println(result)
	result = Includes([]uint8{1, 2, 3, 4}, 2)
	fmt.Println(result)
}

// Includes Constraint operators
func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}
