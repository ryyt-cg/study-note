# Advanced Concurrency

## Review

Goroutine are independently executing functions in the same address space.

```go
go f()
go g(1,2)
```

Channels are typed values that allow goroutines to synchronize and exchange information

```go
c := make(chan int)  // create un-buffer int channel
go func() {
	c <- 3          // send int 3 to the channel
}

n := <-c            // n receive int value from the channel
```

Build long-live system, use select statement

```go
select {
    case xc <- x;
        // sent x on xc 
    case y := <- yc
        // received y from yc
}
```