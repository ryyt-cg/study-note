package main

import "fmt"

// The provided code is written in Go and demonstrates the concept of zero values
// in the language. Zero values are default values assigned to variables that are
// declared without an explicit initial value.
func main() {
	var n int
	var n2 int16
	var f float32
	var b bool
	var s string
	var sl []byte

	fmt.Printf("n - type: %T value: %#v\n", n, n)
	fmt.Printf("n2 - type: %T value: %#v\n", n2, n2)
	fmt.Printf("f - type: %T value: %#v\n", f, f)
	fmt.Printf("b - type: %T value: %#v\n", b, b)
	fmt.Printf("s - type: %T value: %#v\n", s, s)
	fmt.Printf("sl - type: %T value: %#v\n", sl, sl)
}
