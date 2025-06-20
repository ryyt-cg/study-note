package main

import (
	"fmt"
	"math/rand"
	"time"
)

// boring function is a generator of messages and sleeps for a random duration.
// It never ends
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	go boring("boring!")
	fmt.Println("I'm listening.")

	// main function sleeps for 2 seconds and exits
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring; I'm leaving.")
}
