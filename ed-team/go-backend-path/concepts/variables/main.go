package main

import "fmt"

func main() {
	// Variables creation (Static)
	var apple string
	var banana string
	var orange string
	apple = "Apple"
	banana = "Banana"
	orange = "Orange"
	fmt.Println(apple, banana, orange)
	// Variables creation and assignment (Static)
	var pineapple string = "Pineapple"
	fmt.Println(pineapple)
	// Variables creation Static <aggroupment>
	var (
		lemon string = "Lemon"
		peach string = "Peach"
	)
	fmt.Println(lemon, peach)
	// Variables creation and assignment in single line
	var steak, chicken, pork string = "Steak", "Chicken", "Pork"
	fmt.Println(steak, chicken, pork)
	// Variables creation and assignment Dynamic (not recommended)
	fish, number := "Fish", 2
	fmt.Println(fish, number)
	// Variables creation and assignment Dynamic (recommended)
	bird, object := "Bird", "Object"
	fmt.Println(bird, object)
	// You can update values but not types
	// You cannot reasign with the short dynamic operator single values that exists, but news combined:
	apple, integer := "apple", "integer"
	fmt.Println(apple, integer)
}
