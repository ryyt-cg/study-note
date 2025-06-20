package main

import (
	"fmt"
	"sync"
)

var (
	wg      = sync.WaitGroup{}
	counter = 0
	mu      = sync.RWMutex{}
)

func main() {
	//runtime.GOMAXPROCS(100)
	// https://rxmarbles.com/
	for i := 0; i < 10; i++ {
		wg.Add(2)
		mu.RLock()
		go sayHello()
		mu.Lock()
		go increment()
	}

	wg.Wait()
}

func sayHello() {
	defer wg.Done()
	fmt.Printf("Hello #%v\n", counter)
	mu.RUnlock()
}

func increment() {
	counter++
	mu.Unlock()
	wg.Done()
}
