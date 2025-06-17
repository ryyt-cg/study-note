package main

import (
	"fmt"
	"sync"
)

// The output of this program is: 10 times "hello".  The reason is that the goroutine is capturing a copy of the variable msg.
func main() {
	wg := sync.WaitGroup{}
	msg := "Hello"
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(msg string) {
			defer wg.Done()
			fmt.Println(msg)
		}(msg)
	}

	msg = "goodbye"
	wg.Wait()
}
