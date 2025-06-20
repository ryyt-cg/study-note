package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func someFuncWg(num string) {
	defer wg.Done()
	fmt.Println(num)
}

func main() {
	wg.Add(1)
	go someFuncWg("1")

	// continue to do other things.
	fmt.Printf("Greet: %v\n", "Hi")
	wg.Wait()
}
