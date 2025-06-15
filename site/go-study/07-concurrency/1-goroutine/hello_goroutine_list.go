package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	Wg = sync.WaitGroup{}
)

func slow_hello() {
	defer Wg.Done()
	fmt.Println("Hello world goroutine")
	time.Sleep(1 * time.Second)
}

func main() {
	start := time.Now()

	Wg.Add(1000000)
	for i := 0; i < 1000000; i++ {
		go slow_hello()
	}
	Wg.Wait()
	// sleep 1 sec to give hello routine enough time to execute
	//time.Sleep(1 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Processes took %s\n", elapsed)

	fmt.Println("main function")
}
