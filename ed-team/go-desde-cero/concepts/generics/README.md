# Generics

With generics, you can declare and use functions or types that are written to work with any of a set of types provided by calling code.

## No Generic Functions

We can declare a function that receives as `any` type parameters, but beware of the possibilities with an any type:

```Go
func printItems(items ...any) {
	for _, item := range items {
		fmt.Println(item)
	}
}
```

## Type Parameters And Constraints

We can use parameter types to create generics into a function, exists different parameter types, the first one is **Arbitrary Type** constraint

```Go
func sum[T int](nums ...T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}
	return sum
}
```

The other constraint parameter type is **Type with Union Elements**:

```Go
func substr[T int | float32](nums ...T) T {
	var sub T
	for _, num := range nums {
		sub -= num
	}
	return sub
}
```

There is another parameter type constraint, is **Approximation Types**:

```Go
func mult[T ~int | float32](nums ...T) T {
	var mult T
	for index, num := range nums {
		if index == 0 {
			mult = num
		} else {
			mult = mult * num
		}
	}
	return mult
}
```

It is recommended by the course to use the constraints golang package:

```shell
go get golang.org/x/exp/constraints
```

and use as this:

```Go
func div[T constraints.Integer](num ...T) T {
	var div T
	for index, num := range num {
		if index == 0 {
			div = num
		} else {
			div = div / num
		}
	}
	return div
}
```

this package uses union types and approximation Types combination to create powerful parameter types.

## Constraints Operators

Constraints operators help us to compare different types:

```Go
func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}
```

## Generic Types

We can create interfaces (that are a structure type that let us define some types), where we can apply the parameter types with all the constraints mentioned:

```Go
type Product[T uint | string] struct {
	ID          T
	Description string
	Price       float64
}

product1 := Product[uint]{1, "First Product Ever", 3.14}
product2 := Product[string]{"2e", "Second Product Ever", 6.28}
```
