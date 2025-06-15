package main

import (
	"calculator"
	"fmt"
	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Printf("reverse string: %s\n", reverse.String("Hello, World!"))
	fmt.Printf("Sum of two numbers: %v\n", calculator.Add(7, 5))
}
