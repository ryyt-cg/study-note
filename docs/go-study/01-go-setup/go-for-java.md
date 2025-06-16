### Golang vs Java

When it comes to web development, both Golang and Java have advantages and disadvantages. For web development, here is a
feature-by-feature comparison −

Due to its built-in garbage collector and concurrency support, Golang is faster than Java for web development. The
garbage collector in Java can cause latency issues in large-scale applications, affecting performance.

Scalability − Because Golang is so scalable, it's an excellent choice for large-scale web applications. Scalability in
Java is determined by the framework and the developer's ability to manage the application's resources.

Simplicity − Golang is known for its simplicity and ease of use, making it an excellent choice for new programmers.
Because of its extensive libraries and frameworks, Java has a steeper learning curve and can be more complex.

Because of its extensive security features and strong type system, Java is considered more secure than Golang. For
security features, Golang, on the other hand, relies on external libraries.

Libraries and Frameworks − Java has a large number of libraries and frameworks that help developers create complex
web applications. Although Golang has a smaller community and fewer libraries and frameworks, the ones that do exist
are highly efficient and well-documented.

Programming language Maturity

Programmer Community

Tools and Frameworks

### Go and Java have much in common

* C family (imperative, braces)
* Statically typed
* Garbage collected
* Memory safe (nil references, runtime bounds checks)
* Primitive variables are always initialized (zero/nil/false)
* Methods
* Interfaces
* Type assertions (instanceof)
* Reflection

### Go differs from Java in several ways

1. [ ]  Programs compile to machine code. There's no VM.
2. [ ]  Statically linked binaries
3. [ ]  Control over memory layout
4. [ ]  Function values and lexical closures
5. [ ]  Built-in strings (UTF-8)
6. [ ]  Built-in concurrency

### Go intentionally leaves out many features

* No classes
* No constructors
* No inheritance
* No final
* No exceptions
* No annotations

### Why does Go leave out those features?

Clarity is critical.

When reading code, it should be clear what the program will do.

When writing code, it should be clear how to make the program do what you want.

Sometimes this means writing out a loop instead of invoking an obscure function.

(Don't DRY out.)

For more background on design:

Less is exponentially more (Pike, 2012)
Go at Google: Language Design in the Service of Software Engineering (Pike, 2012)

Go looks familiar to Java programmers

Main.java

```
public class Main {
    public static void main(String[] args) {
    System.out.println("Hello, world!");
}
```

hello.go

```go
package main
import "fmt"

func main() {
    fmt.Println("Hello, 世界!")
}

```

Hello, web server

```go
package main

import (
"fmt"
"log"
"net/http"
)

func main() {
    http.HandleFunc("/hello", handleHello)
    fmt.Println("serving on http://localhost:7777/hello")
    log.Fatal(http.ListenAndServe("localhost:7777", nil))
}

func handleHello(w http.ResponseWriter, req *http.Request) {
    log.Println("serving", req.URL)
    fmt.Fprintln(w, "Hello, 世界!")
}
```

Types follow names in declarations.
Exported names are Capitalized. Unexported names are not.

### Goroutines

Goroutines are like lightweight threads.

They start with tiny stacks and resize as needed.

Go programs can have hundreds of thousands of them.

Start a goroutine using the go statement:

go f(args)
The Go runtime schedules goroutines onto OS threads.

Blocked goroutines don't use a thread.
