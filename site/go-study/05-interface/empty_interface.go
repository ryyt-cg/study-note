package main

import "fmt"

// This code demonstrates the dynamic nature of the empty interface in Go. It can
// hold values of any type, and the type can be determined at runtime.
func main() {
	// i is an empty interface
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

// describe prints the type and value of its argument of an empty interface type.
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
