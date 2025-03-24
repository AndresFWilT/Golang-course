package main

import "fmt"

// Constants are better to use them at package scope
const name string = "Andres"

// If are multiple constants use aggroupment
const (
	dbUrl  string = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	dbName string = "postgres"
)

// iota Identifier
const (
	january int = iota + 1
	february
	marz
	april
	may
	june
)

func main() {
	// constants declared assign immediately the value
	const pi float32 = 3.1415926
	fmt.Println(pi)
	// multiple constants in one line
	const height, width int = 600, 400
	fmt.Println(height, width)
	fmt.Println(name, dbUrl, dbName)
	// constants not need to be used to compile
	fmt.Println(january, may)
}
