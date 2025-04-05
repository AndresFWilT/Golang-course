# Testing in Go

Testing is a process where we can validate if a function, component or software accomplish its goal

The best approach is the automatic testing.

## Why is important

Its important because we can automate process. We can prevent errors. We can validate if after a refactor the functionality keeps working.

## In Go

Go has a package named testing, third party packages are no needed.
To execute the test we must follow some notations. Testing files are not compiled in built binary.

### Notation

test files end with: _test.go

```shell
handlerpersons_test.go
```

## Run test

to run test use:

```shell
go test
```

that command will make go executes all the files with `_test.go` and then executes all functions that starts with `TestXXX`

### Run Specific Test Func

```shell
go test -run <SpecificPartOfFuncTestName>
```

or

```shell
go test -run <regularExpression>
```

## Concepts

The package `testing` comes with several Funcs that allow test. In the course is used `T` and `B` Funcs from the **testing** package. T is used for functional testing and B is for benchmarking.
In a background, when the command `go run test` is executed, Go search for `xxx_test.go` and `func TestXxx(t *testing.T) { ... }` functions to execute their tests. Here is an example:

### Usage with T

```Go
package example

func Sum(a, b int) int {
	return a + b
}
```

```Go
package testexample

import "testing"

func TestSum(t *testing.T) {
    wanted := 5
	got := Sum(2, 3)
	if wanted != got {
		t.Errorf("Sum(2, 3) wanted: %v, got: %v", wanted, got)
    }
}

```

In general, is not a good practice to execute logs inside test, but we can use the:

```Go
t.Logf("format", value)
```

from the T package to log. To watch all that logs use:

```shell
go test -v
```

also `-run` flag can be added:

```shell
go test -v run <param>
```

**Remember** use only `t.Log()` to debug the test, and the combination `t.Errorf()` `t.Fail()` to make easier the trace when a test fails. **When the test fails, the debug logs will be shown**

### Usage with B

The `B` Func from the `testing` package, will let us test Benchmarks of our functions. It will test performance! showing a table with the number of operations and the duration of each execution in almost one second. It's important to do when we need to test the performance of our application.

look at the following example:

```Go
package main

func fibonacciFor(n int) int {
	a, b := 0, 1
	if n == 0 {
		return a
	}

	for i := 2; i <= n; i++ {
		c := a + b
		a, b = b, c
	}

	return b
}

func fibonacciRec(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRec(n-1) + fibonacciRec(n-2)
}

var fibonacciData = []int{0, 1}

func fibonacciRecMem(n int) int {
	if n == 0 {
		return fibonacciData[0]
	}

	if len(fibonacciData) >= n+1 {
		return fibonacciData[n]
	}

	newData := fibonacciRecMem(n-1) + fibonacciRecMem(n-2)
	fibonacciData = append(fibonacciData, newData)

	return newData
}
```

We create the functional test, but also the benchmarks as follows:

```Go
package main

import "testing"

const (
	fiboSequence int = 12
	wanted       int = 144
)

func TestFibonacciFor(t *testing.T) {
	got := fibonacciFor(fiboSequence)
	t.Logf("fibonacciFor(%d) = %d, wanted: %v", fiboSequence, got, wanted)

	if got != wanted {
		t.Errorf("fibonacciFor(%v) = %d; want %d", fiboSequence, got, wanted)
	}
}

func TestFibonacciRec(t *testing.T) {
	got := fibonacciRec(fiboSequence)
	t.Logf("fibonacciRec(%d) = %d, wanted: %v", fiboSequence, got, wanted)

	if got != wanted {
		t.Errorf("fibonacciRec(%v) = %d; want %d", fiboSequence, got, wanted)
	}
}

func TestFibonacciRecMem(t *testing.T) {
	got := fibonacciRecMem(fiboSequence)
	t.Logf("fibonacciRecMem(%d) = %d, wanted: %v", fiboSequence, got, wanted)

	if got != wanted {
		t.Errorf("fibonacciRecMec(%v) = %d; want %d", fiboSequence, got, wanted)
	}
}

func BenchmarkFibonacciFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciFor(30)
	}
}

func BenchmarkFibonacciRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciRec(30)
	}
}

func BenchmarkFibonacciRecMem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciRecMem(30)
	}
}
```

To create benchmark test, the Func must be like this: `func BenchmarkXxx(b *testing.B) { ... }`, and to execute use:

```
go test -bench=.
```

that will execute all the Funcs that starts with **Benchmark**.

### What is going on?:

- `b.N` is the number of iterations the benchmark framework will run.
- The Go testing tool automatically determines how many iterations to use so that the benchmark runs for a long enough time to produce meaningful data.
- You should not set b.N manually — Go increases it until it gets stable results.

If is needed to look at memory allocations, use:

```shell
go test -bench=. -benchmem
```

### TableDriven subtest testing

Is possible to create a set table of tests and then execute them to  verify multiple conditions, also create a `Name` field to the set table in order to test specifically one test. 
Look at the following code:

```Go
package main

func Multiply(a, b int) int {
	return a * b
}
```

To test tha massively in order, we can create a slice of a set test struct as follows:

```Go
package main

import "testing"

type table struct {
	name         string
	a, b, wanted int
}

func TestMultiply(t *testing.T) {
	table := []table{
		{"3x1", 3, 1, 3},
		{"3x2", 3, 2, 6},
		{"3x3", 3, 3, 9},
		{"3x4", 3, 4, 12},
		{"3x5", 3, 5, 15},
		{"3x6", 3, 6, 18},
		{"3x7", 3, 7, 21},
		{"3x8", 3, 8, 24},
		{"3x9", 3, 9, 27},
		{"3x10", 3, 10, 30},
	}
	for _, row := range table {
		t.Run(row.name, func(t *testing.T) {
			got := Multiply(row.a, row.b)
			if got != row.wanted {
				t.Errorf("multiply(%v, %v) = %d, want %d)", row.a, row.b, got, row.wanted)
				t.Fail()
			}
		})
	}
}
```

to execute a specific set of test, use:

```shell
go test -v -run TestMultiply/3x3
```

### Structs Instantiation and Pointers Testing

To test different instances of a struct and compare then, we can simply use the logic operators to compare then:

```Go
objA := SomeStruct{}
objB := SomeStruct{}

if objA != objB {
	t.Errorf("test failed! %v", true)
}
```

but the issue comes when we are comparing pointers from struct instances. For that use and deeply only in testing `reflect.DeepEqual()` function. That function inside will compare the instances completely. In functional code is better to compare fields of each instances instead of using this package.

```Go
objA := &SomeStruct{}
objB := &SomeStruct{}

if reflect.DeepEqual(objA, objB) {
	t.Errorf("test failed! %v", true)
}
```

### Testing APIs

#### With net/http

The net/http package has a `httptest` module that will help use to create test.

As all handlers receives as params a **writer** `http.ResponseWriter` and a **request** `*r.Request` we can use the httptest to create a Mock of those params as follows:

```Go
package test

import (
	"net/http"
	"testing"
	"net/http/httptest"
)

func TestGet(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	// call handler
	Get(w, r)
	
	if w.Code != http.StatusInternalServerError {
		t.Errorf("wanted status code: %v, got: %v", http.StatusOK, w.Code)
    }

	// compare objects
	got := Person{}
	err := json.NewDecoder(w.Body).Decode(&got)
	if err != nil {
		t.Errorf("no se pudo procesar el json: %v", err)
	}

	want := Person{
		Name: "Jhoana",
		Age:  31,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Se esperaba %v, se obtuvo %v", want, got)
	}
}
``` 

#### With Echo

Echo uses the `net/http` package, so it's also necessary to create a Mock instance of the writer and request params using the httptest.

```Go
w := httptest.NewRecorder()
r := httptest.NewRequest(http.MethodGet, "/", nil)
e := echo.New()
ctx := e.NewContext(r, w)
Get(ctx)
```

And else keeps the same

### Integration Test

It's a test that ensures functionality that depends on external services (BD, Mail, Communications).
**It is only an integration if it uses an external service**

```go
func HandlerSaveInDB(db databaseInterface, name string){
	db.Save(name)
}
```

In that case, the db package must be an interface. In that situation there are two options:

- Connect to the DB in the Test, `Integration Test`
- Simulate the connection to the DB in the Test, (**Mock**) `Unit Test`

### Mocks

To create Mocks we can create a struct that has methods that are inherit from the interface (db for example). We can create the mocks we want that implement the interface. So instead of passing the correct `usecase` that implements the interface, we can pass the mock to emulate the functionality instead of saving in the db:

Interface of Port

```Go
type Storage interface {
	Create(person *model.Person) error
	Update(ID int, person *model.Person) error
	Delete(ID int) error
	GetByID(ID int) (model.Person, error)
	GetAll() (model.Persons, error)
}
```

Mock

```Go
package handler

import (
	"errors"
)

type PersonStorageOKMock struct{}

func (psm *PersonStorageOKMock) Create(person *model.Person) error {
	return nil
}
func (psm *PersonStorageOKMock) Update(ID int, person *model.Person) error {
	return nil
}
func (psm *PersonStorageOKMock) Delete(ID int) error {
	return nil
}
func (psm *PersonStorageOKMock) GetByID(ID int) (model.Person, error) {
	return model.Person{}, nil
}
func (psm *PersonStorageOKMock) GetAll() (model.Persons, error) {
	return nil, nil
}

type PersonStorageWrongMock struct{}

func (psm *PersonStorageWrongMock) Create(person *model.Person) error {
	return errors.New("error")
}
func (psm *PersonStorageWrongMock) Update(ID int, person *model.Person) error {
	return errors.New("error")
}
func (psm *PersonStorageWrongMock) Delete(ID int) error {
	return errors.New("error")
}
func (psm *PersonStorageWrongMock) GetByID(ID int) (model.Person, error) {
	return model.Person{}, errors.New("error")
}
func (psm *PersonStorageWrongMock) GetAll() (model.Persons, error) {
	return nil, errors.New("error")
}
```

test

```Go
func TestCreate_wrong_storage(t *testing.T) {
	data := []byte(`{"name": "Alexys", "age": 40, "communities": []}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))
	r.Header.Set("Content-Type", "application/json")
	e := echo.New()
	ctx := e.NewContext(r, w)

	p := newPerson(&PersonStorageWrongMock{})
	err := p.create(ctx)
	if err != nil {
		t.Errorf("no se esperaba error, se obtuvo %v", err)
	}

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Código de estado, se esperaba %d, se obtuvo %d", http.StatusInternalServerError, w.Code)
	}

	resp := responsePerson{}
	err = json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		t.Errorf("no se pudo leer el cuerpo: %v", err)
	}

	wantMessage := "Hubo un problema al crear la persona"
	if resp.Message != wantMessage {
		t.Errorf("mensaje de error equivocado: se esperaba %q, se obtuvo %q", wantMessage, resp.Message)
	}
}
```

## Coverage

to look at test coverage use:

```shell
go test -coverprofile=profile.out
```

To a more readable output:

```shell
go tool cover -func=profile.out
```

To a html output:

```shell
go tool cover -html=profile.out
```