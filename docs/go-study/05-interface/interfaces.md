# Interfaces
When to use interfaces:
* Polymorphism
* Behavior
* Mocking
* Generics - please use generics if you are using Go 1.18 or later.




## The Empty Interface
* The interface type that specifies zero methods is known as the empty interface.
* Rob Pike says that empty interface (interface{}) says nothing.

```go
interface{}
```

An empty interface may hold values of any type. Every type implements at least zero methods.

Empty interfaces are used by code that handles values of an unknown type. For example, fmt.Print takes any number of
arguments of type interface{}.

**Example**:  The describe function in the following code accepts an empty interface as an argument. It prints the value and
the type of the argument passed to it. Its application is similar to generics which available in Go 1.18.

[embedmd]:# (./empty_interface.go)

```go
package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

An interface type is defined as a set of method signatures. When a type provides definition for all the methods in the
interface, it is said to implement the interface. It is much similar to the OOP world. Interface specifies what methods
a type should have and the type decides how to implement these methods.

A value of interface type can hold any value that implements those methods.
