package main

import "fmt"

func main() {
	// Create a slice of strings.
	// Initialize the 100th element with an empty string.
	slice := []string{99: "99"}
	fmt.Println(slice)

	// Create a nil slice of integers.
	var nilSlice []int
	fmt.Println(nilSlice)
}
