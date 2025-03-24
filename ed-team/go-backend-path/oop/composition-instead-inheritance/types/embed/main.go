package main

import "fmt"

type Person struct {
	Name string
	Age  uint8
}

func (p Person) Greet() {
	fmt.Printf("Hello, %s!\n", p.Name)
}

func NewPerson(name string, age uint8) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

// Employee embedding with Person
type Employee struct {
	Person
	Salary float32
}

func NewEmployee(name string, age uint8, salary float32) Employee {
	return Employee{
		Person: NewPerson(name, age),
		Salary: salary,
	}
}

func (e Employee) Payroll() {
	fmt.Printf("Payroll called for %s with age %d is: %+v\n", e.Name, e.Age, e.Salary*0.90)
}

func main() {
	p := Person{"Maria", 26}
	fmt.Println(p)
	e := NewEmployee("Maria", 26, 27)
	e.Payroll()
}
