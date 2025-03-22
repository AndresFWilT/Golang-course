package main

import "fmt"

func main() {
	// functions as parameter
	nums := []int{1, 12, 23, 98, 71, 29}
	result := filter(nums, func(num int) bool { return num%2 == 0 })
	fmt.Println(result)

	// functions as return statement
	sumTwo := sum(2)
	fmt.Println(sumTwo(5))
}

func filter(nums []int, callback func(int) bool) []int {
	result := make([]int, 0, len(nums))
	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}

	return result
}

func sum(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}
