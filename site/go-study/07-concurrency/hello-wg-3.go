package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for _, salutation := range []string{"hello", "greeting", "good day"} {
		wg.Add(1)

		// proper way is to pass a copy of salutation into the closure
		// so that by the time the goroutine is run it will be operating on the
		// dta from its iteration of the loop:
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}

// the result is depended on goroutine scheduler. My two runs' result:
//good day
//hello
//greeting
//----------
//good day
//greeting
//hello
