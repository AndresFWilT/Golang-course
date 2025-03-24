package main

// format package
import "fmt"

func main() {
	var a bool = true
	var b string = "golang"
	var c int = -2
	var d uint = 3
	// The verbs %T are allocation indicators
	fmt.Printf("type: %T, value: %v \n", a, a)
	fmt.Printf("type: %T, value: %v \n", b, b)
	fmt.Printf("type: %T, value: %v \n", c, c)
	fmt.Printf("type: %T, value: %v \n", d, d)
	const age uint8 = 255
	fmt.Printf("type: %T, value: %v \n", age, age)
	const ageAlias byte = 254
	fmt.Printf("type: %T, value: %v \n", ageAlias, ageAlias)
	// Types for Unicode
	const unicode rune = 'a'
	fmt.Printf("type: %T, value: %v \n", unicode, unicode)
	// You cannot relate different types, for example sum this:
	var al uint8 = 255
	var bl uint16 = 3000
	// cannot: var c := al + bl
	cl := uint16(al) + bl
	fmt.Printf("type: %T, value: %v \n", cl, cl)
	// blank identifier only with =
	_ = uint8(cl)

	// Zero value only applied to vars
	var string2 string
	var bool1 bool
	var uint2 uint8
	fmt.Printf("type: %T, value: %v \n", string2, string2)
	fmt.Printf("type: %T, value: %v \n", bool1, bool1)
	fmt.Printf("type: %T, value: %v \n", uint2, uint2)
}
