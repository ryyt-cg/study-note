package main

import "fmt"

func main() {
	result := addInt(2, 5)
	resultF := addFloat(4.82, 8.94)
	fmt.Printf("result: %v\n", result)
	fmt.Printf("result: %v\n", resultF)
}

func addInt(a int, b int) int {
	return a + b
}

// addFloat is a function that takes two float64 values and returns their sum.
func addFloat(a float64, b float64) float64 {
	return a + b
}
