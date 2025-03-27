# Sync Mutex

A Mutex (Mutual Exclusion) ensures that only 1 goroutine can access a critical section of code at a time. Without mutex and locking, the goroutines might:

- Read/Write the balance at the same time.
- Cause data corruption in race conditions.

Given the main sample in the concurrency race conditions problem:

```Go
package main

import (
	"fmt"
	"sync"
)

var (
	balance = 100
)

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	lock.Lock()
	b := balance
	balance = b + amount
	lock.Unlock()
}

func Balance() int {
	b := balance
	return b
}

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println("Balance = ", Balance())
}

```

explaining:

```Go
func Deposit(amount int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done() // Mark this goroutine as “done” when finished
	lock.Lock()     // Ensure exclusive access to the balance
	b := balance    // Read balance
	balance = b + amount // Write new balance
	lock.Unlock()   // Release the lock so others can proceed
}
```

## Read and Writing Mutex

`RWMutex` is simply a Read-Write Mutex

It allows **MULTIPLE** concurrent Readers (`RLock`), only **ONE** Writer at a time (`Lock`) and **no readers allowed during writing** — write has exclusive access.

| Operation | 	sync.Mutex |	sync.RWMutex |
| --------- | ------------ | --------------- |
| Lock() |	Exclusive |	Exclusive (for writing) |
| Unlock() |	Releases lock |	Releases lock |
| ❌ Allows concurrent reads? |	❌ No |	✅ Yes |
| ✅ Good for read-heavy code? |	❌ Not optimal |	✅ Yes |

So, in the code we can:

Uses RLock() (read lock) — allows other reads to happen simultaneously, but if a write is happening, the read must wait. When multiple Read operations calls happen concurrently, they won’t block each other

```Go
lock.RLock()
b := balance
lock.RUnlock()
```

Full example:

```Go
package main

import (
	"fmt"
	"sync"
)

var (
	balance = 100
)

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	lock.Lock()
	b := balance
	balance = b + amount
	lock.Unlock()
}

func Balance(lock *sync.RWMutex) int {
	lock.RLock()
	b := balance
	lock.Unlock()
	return b
}

func main() {
	var wg sync.WaitGroup
	var lock sync.RWMutex
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println("Balance = ", Balance(&lock))
}
```


### Summary

|Feature | 	Why it's used                                              |
| -- |-------------------------------------------------------------|
|sync.RWMutex | 	Better for read-heavy concurrent access                    |
|RLock / RUnlock | 	Allow multiple goroutines to read safely                   |
|Lock / Unlock | 	Prevent other access while writing                         |
|Locking reads? | 	✅ Needed to avoid race conditions during concurrent access |