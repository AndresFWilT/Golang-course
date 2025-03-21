package main

import "fmt"

type Request struct {
	// it's better to upper case the fields of the struct in order to make them exportable or accessible from other packages
	RequestId string
}

func main() {
	request := Request{
		RequestId: "1234",
	}
	fmt.Printf("%+v\n", request)
	// pointers
	requestPointer := &request
	(*requestPointer).RequestId = "54321"
	fmt.Printf("%+v\n", requestPointer)
	fmt.Printf("%+v\n", request)
}
