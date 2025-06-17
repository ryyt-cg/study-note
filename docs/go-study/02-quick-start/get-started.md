# Get Started

Create a hello directory

```shell
mkdir hello
cd hello
```

Enable dependency tracking

```shell
go mod init example/hello
go: creating new go.mod: module example/hello 
```

Add main.go to the project

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go")
}
```

Run your code to see the greeting

```shell
go run main.gp
```

## Call code in an external package

When you need your code to do something that might have implemented by someone else. Let's use rsc.io/quote package

import rsc.io/quote package

```go
package main

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
```

Update your dependency tracking<br/>
Run your code

```shell
go mod tidy
go run main.go

Don't communicate by sharing memory, share memory by communicating.
```