# Using Waitgroup

Waitgroup syntax overview

```go
var wg = sync.Waitgroup{}
wg.Add(<int>)               // increase counter by int
wg.Done()                   // decrement counter
wg.Wait()                   // BLOCKING - waiting other goroutines to finish or wg decrement to zero then it continues
```