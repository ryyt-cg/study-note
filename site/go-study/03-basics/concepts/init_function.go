package main

import "fmt"

// sample functions - https://tutorialedge.net/golang/the-go-init-function/
// The init function is a special function in Go that is executed before the main function.

func init() {
	fmt.Println("1st init function starts")
}

func init() {
	fmt.Println("2nd init function starts")
}

func main() {
	fmt.Println("main function starts")
}
