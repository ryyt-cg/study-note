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
