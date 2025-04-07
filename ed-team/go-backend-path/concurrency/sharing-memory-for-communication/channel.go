package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

var urls = []string{
	"http://localhost:1234?duration=3s",
	"http://localhost:1234?duration=1s",
	"http://localhost:1234?duration=5s",
}

func main() {
	fetchConcurrent(urls)
}

func fetchSequential(urls []string) {
	for _, url := range urls {
		fetch(url)
	}
}

func fetchConcurrent(urls []string) {
	signal := make(chan struct{})

	for _, url := range urls {
		go func(u string) {
			fetch(u)
			signal <- struct{}{}
		}(url)
	}

	<-signal
	<-signal
	<-signal
}

func fetch(url string) {
	var values = strings.Split(url, "?duration=")
	var seconds = strings.Split(values[1], "s")
	second, err := strconv.Atoi(seconds[0])
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Duration(second) * time.Second)
	log.Println(url, ": ", 200)
}
