# Pointers in Go

Pointers are variables that store the memory allocation of a value

## Address of Operator

The address of operator is used when you want to know the allocation memory of a variable:

```Go
var variable string = "some value"
var pointer *string
pointer = &variable
```

## Dereference Operator

It is used to get the value allocated in the pointer referenced to a value

```Go
var variable string = "some value"
var pointer *string
pointer = &variable
var value string = *pointer
```

## Update variables values from the pointer

You can update the variables value from the pointer as follows

```Go
var variable string = "some value"
var pointer *string
pointer = &variable
*pointer = "Another value"
```
