package main

import "fmt"

type Course struct {
	name string
}

func (c Course) Print() {
	fmt.Printf("Course name: %+v", c.name)
}

// alias declaration
type myAlias = Course

// types definition
type NewCourse Course

type newBool bool

func (b newBool) String() string {
	if b {
		return "VERDADERO"
	}
	return "FALSO"
}

func main() {
	c := Course{"Golang"}
	c.Print()
	fmt.Printf("Type of the variable instance: %T\n", c)
	c1 := myAlias{"JavaScript"}
	c1.Print()
	fmt.Printf("Type of the variable instance: %T\n", c1)
	c2 := NewCourse{"Python"}
	// The type is NewCourse and we cannot call to the Print Course method.
	fmt.Printf("Type of the variable instance: %T\n", c2)
	// we can create types over defined types
	b := newBool(true)
	fmt.Printf("Value: %+v, Type of the variable instance: %T\n", b.String(), b)
}
