package main

import (
	"fmt"
	"strings"
)

func main() {
	result := sum(2, 3)
	fmt.Println(result)

	// functions with multiple return statements
	text := "Andres"
	lower, upper := transform(text)
	fmt.Println(lower, upper)
}

func sum(a, b int) int {
	return a + b
}

// Recommended only if the functions are short
func transform(text string) (lower string, upper string) {
	lower = strings.ToLower(text)
	upper = strings.ToUpper(text)
	return
}

// if not better use this:
func convert(text string) (string, string) {
	return strings.ToLower(text), strings.ToUpper(text)
}
