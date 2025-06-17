# Array

array - fixed-length data type

Declaring an array set to its zero value

```go
// Declare an integer array of five elements
var array [5]int
```

Declaring an array using an array literal

```go
// Declare an integer array
// Initialize each element with a specific value
// Capacity is determined based on the number of values initialized
array := [...]int{10, 20, 30, 40, 50}
```

Declaring an array initializing specific elements

```go
// Declare an integer array of five elements
// Initialize index 1 and 2 with specific value
// The rest of the elements contain their zero value.
array := [5]int{1:10, 2:20}
```

Accessing array elements

```go
package main

import "fmt"

func main() {
	array := [5]int{10, 20, 30, 40, 50}

	// Change the value at index 2,
	array[2] = 35
	fmt.Println(array)
}

```

Accessing array pointer elements

```go
// Declare an integer pointer array of five elements.
// Initialize index 0 and 1 of the array with integer pointers.
array := [5]*int{0: new(int), 1: new(int)}
```

Assigning one array to another of the same type

```go
// Declare a string array of five elements
var array1 [5]string
fmt.Println(array1)

// Declare  second string array of five elements.
// Initialize the array with colors
array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
fmt.Println(array2)

// Copy the values from array2 into array1
array1 = array2
fmt.Println(array1)

```

Assigning one array of pointers to another

```go
// Declare a string pointer array of three elements.
var array1 [3]*string

// Declare a second string pointer array of three elements.
// Initialize the array with string pointers.
array2 := [3]*string{new(string), new(string), new(string)}

// Add colors to each element
*array2[0] = "Red"
*array2[1] = "Blue"
*array2[2] = "Green"

// Copy the values from array2 into array1. both point to the same memory address
array1 = array2
```

Declaring two-dimensional arrays

```go
// Declare a two dimensional integer array of four elements
// by two elements.
var array [4][2]int

// Use an array literal to declare and initialize a two
// dimensional integer array.
array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}

// Declare and initialize index 1 and 3 of the outer array.
array := [4][2]int{1: {20, 21}, 3: {40, 41}}

// Declare and initialize individual elements of the outer
// and inner array.
array := [4][2]int{1: {0: 20}, 3: {1: 41}}

```

### Passing arrays between functions

Passing an array between functions can be an expensive operation in terms of memory and performance. When you pass
variables between functions, they’re always passed by value. When your variable is an array, this means the entire
array, regardless of its size, is copied and passed to the function.

### Passing a large array by pointer between functions

```go
// Allocate an array of 8 megabytes.
var array [1e6]int

// Pass the address of the array to the function foo.
foo(&array)

// Function foo accepts a pointer to an array of one million integers.
func foo(array *[1e6]int) {
    ...
}
```