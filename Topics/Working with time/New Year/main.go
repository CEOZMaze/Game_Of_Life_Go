package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
 a := [4]int
	nextYear := now.Year() + 1

	firstJanuary := time.Date(nextYear, 1, 1, 0, 0, 0, 0, time.Local)

	// The hidden 'check' function checks your answer:
	check(firstJanuary.Sub(now))
	fmt.Println("")
}
