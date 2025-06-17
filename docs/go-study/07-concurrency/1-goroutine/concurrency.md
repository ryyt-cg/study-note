goroutines - to execute tasks independently, potentially in parallel
channels - for communication, synchronization between goroutines

```go
package main

func main() {
   tasks := getTasks()

   // Process each task.
   for _, task := range tasks {
      process(task)
   }
   
}
```

### WaitGroup

WaitGroup is a great way to wait for a set of concurrent operations to complete when you either don’t care about the
result of the concurrent operation, or you have other means of collecting their results. If neither of those
conditions are true, I suggest you use channels and a select statement instead.

You can think of a WaitGroup like a concurrent-safe counter: calls to Add increment the counter by the integer passed
in, and calls to Done decrements the counter by one. Calls to Wait block until the counter is zero.

### Mutex @ RWMutex

Mutex stands for "mutual exclusion" and is a way to guard critical sections of your program. Mutex creates a share
memory and provides a synchronized access to that memory.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	// Increment
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// Decrement
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}
	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")
}
```

#### channels are

* goroutine-safe
* store and pass values between goroutines
* provide FIFO semantics
* can cause goroutines to block and unblock

#### making channels

buffered channel

```go
ch := make(chan Task, 3)
```

unbuffered channel

```go
ch := make(chan Task)

```

Send and Receive

G1

```go
package main;

func main() {
   ...
   for _, task := range tasks {
      taskCh <- task
   }
}
```

G2

```go
func worker() {
   for {
      task := <-taskCh
      process(task)
   }
}
```
