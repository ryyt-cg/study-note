package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	go someFunc("1")
	// continue to do other things.
	fmt.Printf("Greet: %v\n", "Hi")

	// block 1 second giving enough time to execute someFunc goroutine.
	time.Sleep(1 * time.Second)
}
