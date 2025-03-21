# Get Going with Go

Go was created at Google. Go is a language that:

- Compiles quickly
- Efficient execution
- Ease of programming

Go uses Unicode in UTF-8

Comments:

```Go
//
/* */
```

structure:

```Go
package main

import "fmt"

func main () {
	fmt.Println("Hello World!")
}
```

fmt is a package from the Golang standard library, is "formatted package".

Function syntax:

```Go
func identifier(params...)(return statement)
```

for further details visit [Golang spec](https://go.dev/ref/spec)