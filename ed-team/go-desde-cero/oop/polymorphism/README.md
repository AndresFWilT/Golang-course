# Polymorphism

We can implement polymorphism via interfaces.

an *Interface* is a set of signed methods. In Go is also common to use a simple method via interface. The name of an interface is: **methoder** (must end with `er`)

Values from interfaces can be considered as tuples. Types can satisfy multiple interfaces

In the interfaces composition, we must have the both names of the interfaces to create a new one:

```Go
type Greeter interface{
	Method1()
}

type Byer interface{
	Method2()
}

type GreeterByer interface{
	Greeter
	Byer
}
```

both methods must be disjuncts (method names must be different named)

Also, if we have methods that are pointers, we better use pointers in all the methods that implements an interface (update data), or all the methods that are *normal *

## Empty Interfaces

We can define empty interfaces with 0 structs implementing them, or interfaces with 0 methods:

```Go
package example

import "fmt"

type Exampler interface {
	x()
}

func Something() {
	e := Exampler
	fmt.Printf("%+v, %T",e ,e) // nil and nil
	e.x() // error
}
```

Empty interfaces:

```Go
package main
type Emptier interface {}
```

are useful to handle unknown methods. Any existent type will implement this interface:

```Go
package main

import "fmt"

func wrapper(i interface{}) {
	fmt.Printf("Value: %+v, Type: T%",i,i)
}

func main(){
	wrapper("")
	wrapper(1)
	wrapper(true)
}
```

are really flexible if we don't know the type we are receiving as parameter.

### Type Assertions

We can asser from an empty interface the type of the parameter:

```Go
package sample

import (
	"fmt"
	"strings"
)

func wrapper(i interface{}) {
	value, ok := i.(string) // string assertion
	if ok {
		fmt.Println(strings.ToUpper(value))
	}
}
```

### Type Switch

We can implement a type Switch in order to get multiple assertions of types:

```Go
package main

import (
	"fmt"
	"strings"
)

func wrapper(i interface{}){
	switch value := i.(type) {
	case string:
		fmt.Println(strings.ToUpper(value))
	case int:
		fmt.Println(value*2)
	default:
		fmt.Printf("Value: %+v, Type: %T", value, value)
	}
}
```
