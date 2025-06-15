package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type CustomData interface {
	constraints.Ordered | []byte
}

type User[T CustomData] struct {
	ID   int
	Name string
	Data T
}

// This code demonstrates the power of generics and interfaces in Go, allowing
// for type-safe code reuse. The User struct can be used with any type that
// matches its type constraint, eliminating the need for duplicate code for
// different types.
func main() {
	u := User[int]{
		ID:   0,
		Name: "int",
		Data: 2,
	}
	fmt.Printf("int user: %v\n", u)

	ustring := User[string]{
		ID:   1,
		Name: "string",
		Data: "two",
	}

	fmt.Printf("int user: %v\n", ustring)
}
