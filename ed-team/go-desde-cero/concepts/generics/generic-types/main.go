package main

import "fmt"

type Product[T uint | string] struct {
	ID          T
	Description string
	Price       float64
}

func main() {
	product1 := Product[uint]{1, "First Product Ever", 3.14}
	product2 := Product[string]{"2e", "Second Product Ever", 6.28}
	fmt.Println(product1)
	fmt.Println(product2)
}
