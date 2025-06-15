package main

import (
	"fmt"
	"sync"
)

var wgh = sync.WaitGroup{}

func Hello() {
	fmt.Println("Hello world goroutine")
	wgh.Done()
}

func main() {
	wgh.Add(2)
	go Hello()
	wgh.Wait()

	fmt.Println("main function")
}
