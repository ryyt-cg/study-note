# Channel

Creating a channel in Go requires the use of the built-in function make.

```go
// Unbuffered channel of integers.
unbuffered := make(chan int)

// Buffered channel of strings.
buffered := make(chan string, 10)
```

## Unbuffered Channels

An unbuffered channel is a channel with no capacity to hold any value before it’s received. These types of channels
require both a sending and receiving goroutine to be ready at the same instant before any send or receive operation can
complete. If the two goroutines aren’t ready at the same instant, the channel makes the goroutine that performs its
respective send or receive operation first wait. Synchronization is inherent in the interaction between the send and
receive on the channel. One can’t happen without the other.

## Buffered Channels

A buffered channel is a channel with capacity to hold one or more values before they’re received. These types of
channels don’t force goroutines to be ready at the same instant to perform sends and receives. There are also different
conditions for when a send or receive does block. A receive will block only if there’s no value in the channel to
receive. A send will block only if there’s no available buffer to place the value being sent. This leads to the one big
difference between unbuffered and buffered channels: An unbuffered channel provides a guarantee that an exchange between
two goroutines is performed at the instant the send and receive take place. A buffered channel has no such guarantee.

### Summary

* Concurrency is the independent execution of goroutines.
* Functions are created as goroutines with the keyword go.
* Goroutines are executed within the scope of a logical processor that owns a single operating system thread and run
  queue.
* A race condition is when two or more goroutines attempt to access the same resource.
* Atomic functions and mutexes provide a way to protect against race conditions.
* Channels provide an intrinsic way to safely share data between two goroutines.
* Unbuffered channels provide a guarantee between an exchange of data. Buffered channels do not.