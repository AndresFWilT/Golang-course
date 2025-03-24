package main

import (
	"encapsulation/car"
	"fmt"
)

func main() {
	carInstance := car.New("Mazda", "10 ft", 38000.25)
	carInstance.SetModel("CX-30")
	fmt.Printf("%+v", carInstance)
}
