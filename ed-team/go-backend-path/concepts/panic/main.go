package main

import "fmt"

func main() {
	a, b := 60, 0
	fmt.Printf("Division: %v / %v = %v\n", a, b, division(a, b))
	a, b = 11, 12
	fmt.Printf("Division: %v / %v = %v\n", a, b, division(a, b))
	a, b = 12, 13
	fmt.Printf("Division: %v / %v = %v\n", a, b, division(a, b))
}

func division(dividend, divisor int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic: %v", r)
		}
	}()
	validateDivisor(dividend, divisor)
	return dividend / divisor
}

func validateDivisor(dividend, divisor int) {
	if divisor == 0 {
		panic("divisor is zero")
	}
}
