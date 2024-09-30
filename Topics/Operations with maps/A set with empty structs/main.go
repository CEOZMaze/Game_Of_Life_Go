package main

import "fmt"

func main() {
	// Create a set named 'mySet' with an empty struct as the value type below.
	// The set can have as many keys as you want; you can get creative!
	mySet := map[string]struct{}{
		"teste":  {},
		"teste2": {},
		"teste3": {},
	}
	fmt.Println()

	checkSetTypeValues(mySet) // DO NOT delete this line!
}
