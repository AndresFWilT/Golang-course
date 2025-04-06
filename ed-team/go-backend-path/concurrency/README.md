# Concurrency

Is the ability to handle multiple tasks by interleaving their execution. It means: structuring the application so that multiple independent processes or operations can run at the same time conceptually.

- Deals with multiple tasks making progress simultaneously, not necessarily at the same instant.
- It's about managing multiple tasks logically
- Go emphasizes concurrency with goroutines and channels.

In other hand, `Parallelism` refers to executing tasks simultaneously by leveraging multiple processors or cores. It physically runs more than one operation at the exact same instant.

- Requires hardware support (multiple cores or processors)
- Aims for increased performance through simultaneous execution
- In Go, parallelism is achieved by having multiple goroutines execute simultaneously on separate cores (configured by `GOMAXPROCS`)

| ASPECT | CONCURRENCY                                 | PARALLELISM                                                   |
| -------|---------------------------------------------|---------------------------------------------------------------|
| CONCEPT | Tasks run independently, managed logically  | Tasks run simultaneously, physically                          |
| HARDWARE | Can be achieved on single-core CPUs         | Requires multi-core CPUs                                      |
| GOAL | Structuring tasks for efficient utilization | Speed up execution by simultaneous use of hardware resources. | 
| Go EXAMPLE | Goroutines and Channels                     | Goroutines with `GOMAXPROCS > 1`                              |

So, **Concurrency** is structuring tasks to handle multiple things efficiently, and **Parallelism** is running multiple tasks simultaneously.

## Goroutines

Goroutines are lightweight threads managed by Go’s runtime. They execute functions or methods concurrently and independently, yet share the same memory space with other goroutines. Because of this shared memory space, communication and synchronization between goroutines is typically managed through channels or synchronization primitives like mutexes.

To create a goroutine:

```Go
go MyFunction()
```

## Fork Join Model

Every Go program always has a Goroutine (main). When we execute the command `go run main.go` we create the goroutine main to execute the main function.

![Fork-Join-Model](/assets/forkjoin-model.png)

