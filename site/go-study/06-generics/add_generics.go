package main

import "fmt"

// This code demonstrates the power of generics in Go, allowing for type-safe
// code reuse. The add function can be used with any type that matches its type
// constraint, eliminating the need for duplicate code for different types.
func main() {
	fmt.Printf("result: %v\n", add(5, 15))
	fmt.Printf("result: %v\n", add(5.9, 8.3))
}

// add is a generic function that takes two values of the same type and returns
// their sum. The type parameter T is constrained to be either int or float64.
func add[T int | float64](a T, b T) T {
	return a + b
}
