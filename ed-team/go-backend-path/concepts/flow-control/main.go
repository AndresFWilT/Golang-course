package main

import "fmt"

func main() {
	// IF
	a, b := 5, 10
	if a < b {
		fmt.Println("a < b")
	} else if a == b {
		fmt.Println("a == b")
	} else {
		fmt.Println("a > b")
	}

	// control flow scope variables
	if c, d := 7, 6; c < d {
		fmt.Println("c < d")
	} else if c == d {
		fmt.Println("c == d")
	} else {
		fmt.Println("c > d")
	}

	var expression string = "value"

	// Switch
	switch expression {
	case "":
	case "1":
	case "232":
		fmt.Println("Ogh")
	default:
		fmt.Println("expression is invalid")
	}

	// Grouping Switch
	/*
		switch {
		case expression == "" || expression == "1":
			fmt.Println("Amazing")
		case "232":
			fmt.Println("Ogh")
		default:
			fmt.Println("expression is invalid")
		}

	*/
}
