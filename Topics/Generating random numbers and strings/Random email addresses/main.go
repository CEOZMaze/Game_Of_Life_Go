package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	var seed int64
	fmt.Scanln(&seed)
	rand.Seed(seed)
	// put your code here
	strinBytes := []byte{}
	sliceString := []string{}
	for x := 0; x < 3; x++ {

		for i := 0; i < 5; i++ {

			out := byte(rand.Intn(26) + 97)
			strinBytes = append(strinBytes, out)

		}
		sliceString = append(sliceString, string(strinBytes))
		strinBytes = []byte{}
		outputString := strings.Join(sliceString, "")
		fmt.Printf("%s@fantasy.com\n", outputString)
		sliceString = []string{}
	}

}
