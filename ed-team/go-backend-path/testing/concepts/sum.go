package main

func Sum(a, b int) int {
	return a + b
}

func MultipleSum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
