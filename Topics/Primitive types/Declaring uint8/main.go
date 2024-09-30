package main

import (
	"fmt"
	"reflect"
)

func solve() uint8 {
	// declare an uint8 type variable 'a' with a value of 10 below:
	var a uint8
	a = 10

	return a // do not delete this line!
}

// DO NOT delete or modify the contents of the main function!
func main() {
	a := solve()
	if reflect.TypeOf(a).Name() == "uint8" {
		fmt.Println(a)
	}
}
