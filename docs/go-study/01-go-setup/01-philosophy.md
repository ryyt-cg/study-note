# Why Go?

## Philosophy

* Simplicity - Go is designed to be simple and easy to learn. It has a clean and easy-to-read syntax that makes it easy
  to write and maintain code.
* Static Typing - Go is statically typed, which means that the type of a variable is known at compile time. This helps
  to catch errors early in the development process.
* Compiled Language - Go is a compiled language, which means that it is fast and efficient. It is designed to be fast to
  compile and fast to run, making it a good choice for performance-critical applications.
* Efficient Memory Management - Go has a garbage collector that automatically manages memory, making it easy to write
  memory-safe code without worrying about memory leaks.
* Dependency Management - Go has a built-in package management system called "go modules" that makes it easy to manage
  dependencies and share code with other developers.
* Concurrency - Go has built-in support for concurrency, which makes it easy to write programs that can do multiple
  things at once. This is especially useful for writing web servers and other networked applications.
* Performance - Go is a compiled language, which means that it is fast and efficient. It is designed to be fast to
  compile and fast to run, making it a good choice for performance-critical applications.
* Cloud Native - Go is designed to work well in cloud-native environments, making it a good choice for building
  microservices and other cloud-based applications.
* Active Community - Go has a growing and active community of developers who are constantly working to improve the
  language and its ecosystem.
* Open Source - Go is an open-source language, which means that it is free to use and can be modified and distributed by
  anyone.
* Standard Library - Go has a rich standard library that provides a wide range of functionality, including networking,
  file I/O, and text processing.

## Go vs Java
https://go.dev/talks/2015/go-for-java-programmers.slide#2

Golang is a multi-paradigm and better supports concurrency. While Golang runs faster than Java, Java has more
features and better support. Go has a leaner learning curve as compared to Java. If simplicity and memory management are
your priority, then Go is a good option

| Parameter          | Golang                                                                        | Java                                                 |
|--------------------|-------------------------------------------------------------------------------|------------------------------------------------------|
| Syntax             | has a simpler syntax with fewer keywords                                      | has a more verbose syntax                            |
| Concurrency        | has built-in support for concurrency with goroutine and channel. light-weight | has concurrency like thread.  Much more heavy-weight |
| Compilation        | is a compiled language with fast compilation times                            | is compiled language to be interpreted in JVM        |
| Garbage collection | has a concurrent garbage collector                                            |                                                      |
| Package management | has a built-in package management system called "go modules"                  | use external tools like Maven or Gradle              |
| Performance        | high-performance and low-latency                                              | slower than Golang since it runs on JVMs             |
| Community support  | has a growing and active community                                            | has a mature community                               |


## Go and Java have much in common

* C family (imperative, braces)
* Statically typed
* Garbage collected
* Memory safe (nil references, runtime bounds checks)
* Primitive variables are always initialized (zero/nil/false)
* Methods
* Interfaces
* Type assertions (instanceof)
* Reflection

## Go differs from Java in several ways

1. [ ]  Programs compile to machine code. There's no VM.
2. [ ]  Statically linked binaries
3. [ ]  Control over memory layout
4. [ ]  Function values and lexical closures
5. [ ]  Built-in strings (UTF-8)
6. [ ]  Built-in concurrency
7. [ ]  Built-in Dependency Management

## Go intentionally leaves out many features

* No classes
* No constructors
* No inheritance
* No final
* No exceptions
* No annotations