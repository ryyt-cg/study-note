package main

import (
	"fmt"
)

func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// close the channel to signal that no more values will be sent
	defer close(ch)
}
func consumer(ch <-chan int, done chan<- bool) {
	sum := 0

	// range over the channel until it is closed
	// instead of using mutexes to synchronize access to shared data
	for value := range ch {
		fmt.Println("Received:", value)
		sum = sum + value
	}

	fmt.Printf("Sum: %d\n", sum)
	done <- true
}
func main() {
	ch := make(chan int)
	done := make(chan bool)
	go producer(ch)
	go consumer(ch, done)

	// block until the consumer is done
	<-done
}
