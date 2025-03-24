# Oriented Object Programming in Go

Is Go an OOP language? **yes and no**.

The OOP has 4 pillars: **Inheritance**, **Abstraction**, **Polymorphism**, **Encapsulation**.

- Abstraction: process to isolate an element from the real world or fiction. In Go we don't have classes but we have structs. The methods are not exclusive from the structures, and can be implemented to maps, funcs, slices, arrays, strings, etc.
- Encapsulation: feature to protect methods or fields from a class (private, public, protected). In Go those identifiers not exists, but we can set which can be exported or not.
- Inheritance: feature that lets a class inherit properties, methods from a father class. **In Go inheritance does not exist**. But with types we can create another types that are integrated similar to inheritance, also has interfaces.
- Polymorphism: feature that lets us implement in different ways a similar feature. In Go we can implement polymorphism media interfaces.