# Structures

In Go no exist classes but structures; structures are simply fields collection. We c
an implement similarities with classes right there.

It's interesting, how we can define methods in a struct, look at the following sample:

```Go
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
```

The method `ListClasses` belongs to the struct Course even if it is defined outside the struct. To call it we can:

```Go
course1 := Course{"Go from 0", 12.34, false, []uint{12, 25, 89}, map[uint]string{
		1: "Introduction", 2: "OOP", 3: "Data Bases", 4: "APIs", 5: "SSR", 6: "Testing", 7: "Concurrency",
	}}

course1.ListClasses()
```

So, the struct Course has a method: a method belongs to a struct and basically is a function with a receiver. The receiver is this:

```
func (c Course) Function() {}
```

The `(c Course)` is the value receiver, which is a copy of the original struct not the original:

- c Receiver name
- Course Receiver type (It's not a pointer but a copy)

That means, that if we want to make changes to the original struct instance, we must do something like this:

```Go
func (c *Course) Function(){
	// amazing stuff updates
}
```
in Go is not a good practice to use something like this, or self to refer to a method.

We can declare those methods that belong to a struct even in other file. but must be inside the same package.

It's better to not mix pointer methods with non pointer methods related to a struct or interface, instead use only one type.