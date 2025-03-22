package main

import "fmt"

func main() {
	printItems("Hola", "Hello", "Bon jour")
	printItems(1, 2, 3, 4, 5)
}

// generic function
func printItems(items ...any) {
	for _, item := range items {
		fmt.Println(item)
	}
}
