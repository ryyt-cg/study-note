# Effective Concurrency

#### Goroutine

Goroutine is an execution context that is managed by the Go runtime ( as opposed to a thread that is managed by the
operating system). A goroutine usually has a much smaller startup overhead than operating system thread.

```go
go f()
go g(i, j)
go func() {
	...
}()

go func(i, j int) {
	
}(1, 2)
```

The number of operating system threads used by the runtime is equal to the number of processors/cores on the
platform (unless you change this by setting the **GOMAXPROCS** environment variable or by calling runtime.**GOMAXPROCS**
function). A goroutine starts with a small task. After version 1.19, goroutine size is about 2K.

```go
package main

import (
	"fmt"
	"time"
)

func f() {
  fmt.Println("Hello from goroutine")
}
func main() {
  go f()
  fmt.Println("Hello from main")
  time.Sleep(100)
}

```

A race condition occurs when two or more threads can access shared data and they try to change it at the same time. 