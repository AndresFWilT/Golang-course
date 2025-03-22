# Functions in Go

Go allow to create functions with value params or reference params

## Value Params

```Go
func main() {
	sayHi("Omg")
}

func sayHi(text string){
	fmt.Println("Hello: ", text)
}
```

Go can infer the type from the params by group as follows

```Go
func somethingAmazing(textA, textB string, numA, numB. numC uint8) {
	// some amazing logic
}

```

## Reference Params

If needed and i think in most of the cases we want to update some values in the function we send:

```Go
func main() {
	someText := "text"
	toUpperCase(&someText)
}

func toUpperCase(someText *string) {
	*someText = strings.toUpper(*someText)
}
```

## With Return Statement

In Go the functions can return multiple parameters as this:

```Go
func convert(text string) (string, string) {
	return strings.ToLower(text), strings.ToUpper(text)
}
```

## Functions as Parameter or as Return Statement

You can set as parameter a function to another function or return it as statement:

```Go
func main() {
    nums := []int{1, 12, 23, 98, 71, 29}
    result := filter(nums, func(num int) bool { return num%2 == 0 })
    fmt.Println(result)
}

func filter(nums []int, callback func(int) bool) []int {
	result := make([]int, 0, len(nums))
	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}

	return result
}
```

return statement

```Go
func main() {
    sumTwo := sum(2)
    fmt.Println(sumTwo(5))
}

func sum(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}
```
