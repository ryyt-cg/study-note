package main

import (
	"fmt"
	"sync"
)

// The output of this program is unpredictable.  The reason is that the goroutine is capturing  a reference of the variable msg.
// The value of msg is changed to "goodbye" before all goroutines completed.
func main() {
	wg := sync.WaitGroup{}
	msg := "Hello"
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(msg *string) {
			defer wg.Done()
			fmt.Println(*msg)
		}(&msg)
	}

	msg = "goodbye"
	wg.Wait()
}
