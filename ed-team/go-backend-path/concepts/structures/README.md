# Structures

Structure will help to store a collection of fields, similar to classes.
It's important to make the structures global to the scope of the package, and also make the fields we want to make exportable with uppercase:

```Go
package main

import (
	"encoding/json"
	"fmt"
)

// Struct with exported (uppercase) fields
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	age   int    // Unexported (won't be serialized)
}

func main() {
	u := User{Name: "John", Email: "john@example.com", age: 30}

	// Convert struct to JSON
	jsonData, _ := json.Marshal(u)
	fmt.Println(string(jsonData)) // {"name":"John","email":"john@example.com"}
}
```

## Pointers

we can create pointers from struct reference A to struct reference B:

```Go
    request := Request{
		RequestId: "1234",
	}
	fmt.Printf("%+v\n", request)
	// pointers
	requestPointer := &request
```

we can update the values from the pointer like this:

```Go
1. requestPointer.RequestId = "54321"
2. (*requestPointer).RequestId = "54321"
```