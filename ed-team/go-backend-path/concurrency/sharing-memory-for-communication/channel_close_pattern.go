package main

import "fmt"

func main() {
	number := make(chan int, 3)
	signal := make(chan struct{})
	go receive(signal, number)
	send(number)

	<-signal
}

func send(number chan<- int) {
	number <- 1
	number <- 2
	number <- 3
	number <- 4
	number <- 5
	close(number)
}

func receive(signal chan struct{}, number <-chan int) {
	for v := range number {
		fmt.Println(v)
	}

	signal <- struct{}{}
}
