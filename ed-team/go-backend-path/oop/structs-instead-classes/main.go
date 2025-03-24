package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func (c Course) ListClasses() {
	text := "The classes in the Course '" + c.Name + "' are: "
	for _, class := range c.Classes {
		text += class + ",\t"
	}
	fmt.Println(text[:len(text)-2])
}

func main() {
	course1 := Course{"Go from 0", 12.34, false, []uint{12, 25, 89}, map[uint]string{
		1: "Introduction", 2: "OOP", 3: "Data Bases", 4: "APIs", 5: "SSR", 6: "Testing", 7: "Concurrency",
	}}
	fmt.Println(course1)
	// dot operator will help us to assign too
	course1.IsFree = true
	course1.ListClasses()
}
