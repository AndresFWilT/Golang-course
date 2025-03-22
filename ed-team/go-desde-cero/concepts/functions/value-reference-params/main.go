package main

import (
	"fmt"
	"strings"
)

func main() {
	variable := say("Amigues Do Alma")
	fmt.Println(variable)

	text := "amigues"
	toUpperCase(text)
	fmt.Println(text)

	text = "Amigues"
	toLowerCase(&text)
	fmt.Println(text)
}

func say(name string) string {
	fmt.Printf("Hello, %s!\n", name)
	return name + "a"
}

// the parameters of a func is a copy of the original variable so that won't update the original
func toUpperCase(text string) {
	text = strings.ToUpper(text)
}

func toLowerCase(text *string) {
	*text = strings.ToLower(*text)
}
