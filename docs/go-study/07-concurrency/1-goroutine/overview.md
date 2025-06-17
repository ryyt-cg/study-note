# Concurrency

## WaitGroups

```go
var wg = sync.Waitgroup{}

wg.Add(<int)
wg.Done() // decrement counter
wg.Wait() // Blocking
```

## Channel

* Channels are used to pass data between / synchronize goroutines, including func main
* Go philosophy: "Don't communicate by sharing memory; share memory by communicating."

Syntax

### unbuffered channel

```go
ch := make(chan string) // unbuffered string channel

ch <- myData    // BLOCKING when no process receives the message
myVar <- ch     // BLOCKING when no message to send

close(ch)       // optional, no longer receive messages
```

### buffered string channel

```go
c := make(chan string, 3)

c <- "Hello "
c <- "Earth "
c <- "from Mars "
c <- "from Venus"   // Deadlock occurs when ran out of buffer
```

Select channel

```go
for _, s := range chars {
    select {
    case charChannelOne <- s:
        fmt.Println("Channel One")
    case charChannelTwo <- s:
        fmt.Println("Channel Two")
    case charChannelThree <- s:
        fmt.Println("Channel Three")
    }
}

```

runtime.GOMAXPROCS(0)

## Race Conditions: Mutexes, Condition, and Atomic Variables

### Mutex

Mutually Exclusive -

```go
mutex = sync.Mutex{}
mutex.Lock()
mutex.Unlock()

```

### Conditional variables

```go
var condition = sync.NewCond(&mutex)
condition.Signal()
condition.Broadcast()
condition.Wait()        // BLOCKING
```

### Atomic variables

```go
atomic.AddInt32(&inventory, 100)
```

equivalent to this code block

```go
mutex.Lock()
inventory += 100
mutex.Unlock()
```

## Concurrency Patterns

* Worker Pool
* Fork & Join