# Concurrency

Concurrency in go is worked with goroutines. Goroutines are a lightweight thread safety library that helps us to work with concurrency.

## The withdraw and deposit problem

Imagine this:

You have a bank account with a balance of 1000 dollars.

One goroutine tries to withdraw 500 dollars. 
at the same time another one tries to deposit 200 dollars.

If both goroutines access and update the balance at the same time, and there's no locking/synchronization, the final balance can become inconsistent or incorrect.

### Use of WaitGroup

A WaitGroup let us **wait until a group of goroutines have finished**.
Every time a goroutine is launch, we must do:

```Go
package sample

import (
	"sync"
)

func SomeFunc(counter int, wg *sync.WaitGroup) {
	defer wg.Done()
	// async logic
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go SomeFunc(i*100, &wg)
	}
	wg.Wait()
}
```

we must first create a WaitGroup via ```var wg sync.WaitGroup```, and on each go routine we are going to launch **(before)** we must do: `wg.Add(1)`. Inside the SomeFunc that receives the pointer of sync.WaitGroup, we `defer wg.Done()` that will decrement the delta (counter of the wg). Finally `wg.Wait()` blocks the main thread until the counter is 0.

Visual flow:

| Step |	What Happens |
| ---  | ---------------- |
| wg.Add(1) |	Register a new goroutine to wait for |
| go ... |	Launch the goroutine |
| wg.Done() |	Called at the end of the goroutine (defer) |
| wg.Wait() |	Waits until all wg.Done()s are called |
