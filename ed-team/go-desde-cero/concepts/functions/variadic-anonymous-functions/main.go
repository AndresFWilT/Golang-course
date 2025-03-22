package main

import "fmt"

func main() {
	// variadic functions
	fmt.Println(sum(1, 2, 23, 4, 5))
	// Anonymous funcs
	a := func() {
		fmt.Println("Hello World")
	}
	println(a)
	func(number int) {
		fmt.Println("Good Bye: ", number)
	}(2)
}

// variadic funcs,
func sum(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}
