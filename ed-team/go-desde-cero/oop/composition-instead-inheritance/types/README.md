# Types in Go

We have types in go which allows us to create structs. We can create different kind of types:

- alias declaration
- types definition

sample:

```Go
package sample

type Course struct {
	name string
}

// Alias declaration
type myCourse = Course
// Types definition
type NewCourse Course
```

In the first case, the alias declaration has the same methods and fields than the original type, but in the types definition, only the fields.

## Embed

We can embed types from others as this:

```Go
package sample

type Person struct {
	name string
	age uint
}

type Employee struct {
	Person
	salary float32
}
```

if we instantiate a new `Employee`, we can access to the Person fields too

## We cannot override in Go

We can do something like this and having in mind the last sample:

```Go
package sample

import "fmt"

func (p Person) Greet(){
	fmt.Println("Hello from Person")
}

func (e Employee) Greet(){
	fmt.Println("Hello from Employee")
}
```

if we execute this:

```Go
e := Employee{"Andres", 26, 1234}
e.Greet()
```

it will print the `Hello from Person` text, but we can also access to the Greet method from Person like this:

```Go
e.Person.Greet()
```

So, that is not override.

### > [!WARNING]  
> If we embed multiple types into another, we can get conflicts:
> 
```Go
package sample

type Human struct {
	age uint
}

type Person struct {
name string
age uint
}

type Employee struct {
Person
Human
salary float32
}
```
> We can get conflict with the field age from both types (Human and Person), like multiple inheritance. To access to one field or other we can: `e.Person.age` or `e.Human.age`
