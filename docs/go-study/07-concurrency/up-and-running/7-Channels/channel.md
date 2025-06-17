# Using Channels

* Channels are used to pass data between / synchronize goroutines, including func main
* Go philosophy:
    * "Don't communicate by sharing memory, share memory by communicating."

## Channel Syntax Overview

```go
ch := make(chan string)
ch <- myData            // (NO BLOCKING)
ch <- myData2           // (BLOCKING) wait until datum in the channel is pulled
myVar <- ch             // (BLOCKING)

close(ch)               // optional - no data can't be sent to this channel
```