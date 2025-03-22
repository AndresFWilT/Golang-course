package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type MyInt int

func main() {
	result := sum[int](1, 2, 3)
	fmt.Println(result)
	resultSubs := substr[float32](1.2, 3.3, 24.5)
	fmt.Println(resultSubs)
	var num1 MyInt = 1
	var num2 MyInt = 2
	resultMult := mult[MyInt](num1, num2)
	fmt.Println(resultMult)
	resultDiv := div[uint64](256, 2)
	fmt.Println(resultDiv)
}

// parameter type with arbitrary type constraint
func sum[T int](nums ...T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}
	return sum
}

// parameter type with union elements constraint
func substr[T int | float32](nums ...T) T {
	var sub T
	for _, num := range nums {
		sub -= num
	}
	return sub
}

// parameter type with approximation types
// The accent makes that all types of int and its derivatives can be passed as generics type for this function
func mult[T ~int | float32](nums ...T) T {
	var mult T
	for index, num := range nums {
		if index == 0 {
			mult = num
		} else {
			mult = mult * num
		}
	}
	return mult
}

// Number You can create an interface which is a type of structure where we can define all necessary constraints

func div[T constraints.Integer](num ...T) T {
	var div T
	for index, num := range num {
		if index == 0 {
			div = num
		} else {
			div = div / num
		}
	}
	return div
}
