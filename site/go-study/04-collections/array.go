package main

import "fmt"

func main() {
	// Declare a string array of five elements
	var array1 [5]string
	fmt.Println(array1)

	// Declare a second string array of five elements.
	// Initialize the array with colors
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println(array2)

	// Copy the values from array2 into array1
	array1 = array2
	fmt.Println(array1)
}
