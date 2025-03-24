package main

import "fmt"

func main() {
	music := make(map[string]string)
	music["Guitar"] = "Gibson"
	music["Magic"] = "Magician"
	fmt.Println(music)
	// literal maps
	tech := map[string]string{
		"Guitar": "Guitar",
		"Magic":  "Magic",
	}
	fmt.Println(tech)

	// deletion
	delete(music, "Guitar")
	fmt.Println(music)
	_, ok := music["Guitar"]
	fmt.Println(ok)
}
