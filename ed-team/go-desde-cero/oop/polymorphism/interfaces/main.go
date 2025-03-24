package main

import "fmt"

type Greeter interface {
	Greet() string
}

type Byer interface {
	Bye()
}

type GreeterByer interface {
	Greeter
	Byer
}

func All(gbs ...GreeterByer) {
	for _, gb := range gbs {
		gb.Greet()
		gb.Bye()
	}
}

type Person struct {
	Name string
}

func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

func (p Person) Bye() {
	fmt.Println("See ya, i am: ", p.Name)
}

type Text string

func (t Text) Greet() string {
	return string("Hello, this is a text: " + t)
}

func (t Text) Bye() {
	fmt.Println("Bye: " + t)
}

func GreetAll(gs ...Greeter) {
	for _, g := range gs {
		g.Greet()
		fmt.Printf("Value: %+v\t Type: %T\n", g, g)
	}
}

func main() {
	g := Person{"John"}
	fmt.Println(g.Greet())
	var t Text = "Hello, world!" // cast to Text
	fmt.Println(t.Greet())
	All(g, t)
}
