# [Advance Concurrency](https://www.youtube.com/watch?v=QDDwwePbDtw)

## Recap
Concurrency is built into the language rather than being a library.  Therefore, it's easy to deal with.

### Goroutines and Channels
Goroutines are independently executing functions in the same address space.
````go
go f()
go g(1,2)
````

Channels are typed values that allow goroutines to synchronize and exchange information.
```go
c:= make(chan int)
go func() { c <- 1 }()  // send 1 to c
n := <-c // receive from c
```

## It's easy to go, but how to stop?
- Long-lived programs need to clean up
- Let's look at how to write programs that handle communication, periodic events, and cancellation.
- The core is Go's select statement: a switch for communication.

```go
select {
case xc <- x:
    fmt.Println("send x to xc")
case y := <-yc:
    fmt.Println("received y from yc", y)		
}
```

