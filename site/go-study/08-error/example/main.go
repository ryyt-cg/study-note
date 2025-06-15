package main

import (
	"errors"
	"fmt"
)

// MyError
// Define a custom error type
type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return e.Msg
}

func main() {
	// Create an instance of the custom error
	myErr := &MyError{"This is a custom error"}

	// Wrap the custom error with additional context
	wrappedErr := fmt.Errorf("an error occurred: %w", myErr)

	// Use errors.Is to check if the wrapped error is a specific error
	if errors.Is(wrappedErr, myErr) {
		fmt.Println("The error is a MyError")
	} else {
		fmt.Println("The error is not a MyError")
	}

	// Use errors.As to extract the custom error from the wrapped error
	var errAsMyError *MyError
	if errors.As(wrappedErr, &errAsMyError) {
		fmt.Println("The error as MyError:", errAsMyError)
	} else {
		fmt.Println("The error could not be converted to MyError")
	}
}
