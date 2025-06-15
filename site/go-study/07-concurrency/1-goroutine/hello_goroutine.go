package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello world goroutine")
}

func main() {
	fmt.Println("Start")
	go sayHello()
	//time.Sleep(100 * time.Millisecond)
	fmt.Println("Exit")
}
