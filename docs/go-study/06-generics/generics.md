# Generics

Generics were proposed a few years ago in Go Forum, and they have finally been accepted into the language. The feature
is in Go 1.18 version.

What do generics really changed in Go?<p/>
Generics allow our functions and data structures to take in several types that are defined in their generic forms.

Let's say you need a function that takes one slice of string and prints it. Then you might write this type of function:

```go
func Print(s []string) {
	for _, v := range s {
	    fmt.Print(v)	
    }
}
```

What if we want to have the slice of integer? You will need to make a new method for that:

```go
func Print(s []int) {
	for _, v := range s {
	    fmt.Print(v)	
    }
}
```

This might see redundant. Therefore, Go 1.18 offers generics. Here is the function might look like this:

```go
func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
}
```

1. We have T, which is the type of any keyword (this keyword is specifically defined as part of a generic, which
   indicates any type)
2. And our parameter, where we have variable s whose type is a slice of T .

---

Problem: create add two numbers function func(a, b)

without generics:
[embedmd]:# (./add_func.go)

```go
package main

import "fmt"

func main() {
	result := addFloat(4.82, 8.94)
	fmt.Printf("result: %v\n", result)
}

func addInt(a int, b int) int {
	return a + b
}

func addFloat(a float64, b float64) float64 {
	return a + b
}
```

applying generics:
[embedmd]:#(./add_generics.go)

```go
package main

import "fmt"

func main() {
	result := add(5, 15)
	fmt.Printf("result: %v\n", result)
}

func add[T int | float64](a T, b T) T {
	return a + b
}
```

The parameter type: T int | float64 is verbose. There are several of the int and float like int32, float32, etc. You
could do this:

```go
type Num interface {
	int | int8 | int16 | float32 | float64
}

func add[T Num](a T, b T) T {
    return a + b
}
```

Better way that Go provides:

Ordered is a constraint that permits any ordered type: any type that supports the operators < <= >= >. If future
releases of Go add new ordered types, this constraint will be modified to include them.

```go
T constraintts.Ordered
```

## Map Generics

```go
package main

import (
	"fmt"
)

type GenericMap[T comparable, V int | string] map[T]V

func main() {
	m := make(GenericMap[string, int])
	m["abc"] = 7
	fmt.Printf("result: %v\n", m)

	m2 := make(GenericMap[string, string])
	m2["abc"] = "seven"
	fmt.Printf("result: %v\n", m2)
}
```

## Struct Generics

```go
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
```
