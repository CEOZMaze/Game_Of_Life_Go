package main

import (
	"errors"
	"fmt"
)

func brokenMessage() error {
	// create the main error here - you can use errors.New or fmt.Errorf
	err := errors.New("ice cream machine is broken")

	// create the wrapped error here using fmt.Errorf and the %w verb
	wrappedErr := fmt.Errorf("error: can't serve you an ice cream main cause of error: %w", err)

	return wrappedErr // Do not delete this line!
}

// DO NOT delete or modify the contents of the main function!
func main() {
	if err := brokenMessage(); err != nil {
		fmt.Println(err)
	}
}
