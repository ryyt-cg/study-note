package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// MapValues
// input: [1, 2, 3]
// output: [2, 4, 6]
// https://pkg.go.dev/golang.org/x/exp/constraints#pkg-types
func MapValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, v := range values {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}

	return newValues
}

func main() {
	result := MapValues([]float64{1.4, 2.7, 3.1}, func(n float64) float64 {
		return n * n
	})

	fmt.Printf("result: %v", result)
}
