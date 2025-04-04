package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/hello-world", handlerHelloWorld)
	println("Server running on http://127.0.0.1:9080")
	http.ListenAndServe(":9080", nil)
}

func handlerHelloWorld(writer http.ResponseWriter, reader *http.Request) {
	fmt.Printf("Hello World from server %v",9)
}
