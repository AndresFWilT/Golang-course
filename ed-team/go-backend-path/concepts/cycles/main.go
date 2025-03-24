package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// This kind of for is useful for websockets and channels when doing concurrency
	j := 0
	for {
		if j == 15 {
			break
		}
		println(j)
		j++
	}

	// for in arrays
	var array []string = []string{"ElementA", "ElementB", "ElementC", "ElementD", "ElementE"}
	for i, v := range array {
		fmt.Printf("Index: %v\t Value: %v\n", i, v)
	}

	// for in slices
	sliceArray := []uint8{3, 6, 9, 12, 15}
	fmt.Println(sliceArray)
	for i := range sliceArray {
		// allocation memory only available via index in array, because the value in the last sample is a copy of the values
		sliceArray[i] *= 2
	}
	fmt.Println(sliceArray)

	// for in map
	foodMap := map[string]string{
		"Pizza":     "Napolitana",
		"Hamburger": "El corral",
	}

	for key, value := range foodMap {
		fmt.Println(key, value)
	}

	// for in strings
	var superString string = "SUPER MEGA STIRNG AMAGINZ"
	for index, value := range superString {
		fmt.Println(index, string(value))
	}
}
