# Golang Fundamentals

## Variables

**DECLARE** a variable to hold a **VALUE** of a certain **TYPE**

### Statical

```Go
var variable int = 2
```

### Dynamic

```Go
variable := 2
```

## Blank Identifier

The blank identifier can be assigned or declared with any value of any type, with the value discarded harmlessly. It represents a write-only value to be used as a placeholder where a variable is needed but the actual value is irrelevant.

## Code Pollution

Declare variables that are not beign used is code polution:

```Go
a := 0
// error!
```

## Zero Value

Zero value is the initialized zero value for their types:

```
var a int
```

## Specification Types

Golang supports natively the following types:

- Boolean
- Numeric
- String
- Array
- Slice
- Struct
- Pointer
- Function
- Interface
- Map
- Channel

## Aggregate Type

Many values together: array, struct, slice

## Composite Type

compound structure <struct>

**The act of constructing a composite type is known as composition**, many different types into one structure which composes all of those different types together into that data structure - It's a DATA STRUCTURE which hols VALUES of many different types.
