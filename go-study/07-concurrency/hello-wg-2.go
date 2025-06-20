package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for _, salutation := range []string{"hello", "greeting", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}

// The result is
// good day
// good day
// good day

// because the goroutines being scheduled may run at any point in time in the future, it is undetermined what values
// will be printed from within the goroutine.
