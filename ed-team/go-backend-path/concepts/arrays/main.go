package main

import "fmt"

func main() {
	// arrays have fixed size
	var flags [3]string
	flags[0] = "A"
	flags[1] = "B"
	flags[2] = "C"

	// literal arrays
	hymns := [3]string{"A", "B", "C"}

	// literal arrays with size inferred
	songs := [...]string{"A", "B", "C"}
	fmt.Println(flags, hymns, songs)
}
