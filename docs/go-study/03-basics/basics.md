# Go Basics

## Blank Identifier

The _ (underscore character) is known as the blank identifier and has many uses within Go. It’s used when you want to
throw away the assignment of a value, including the assignment of an import to its package name, or ignore return values
from a function when you’re only interested in the others.

## init

Each package can provide as many init functions as necessary to be invoked at the beginning of execution time. All the
init functions that are discovered by the compiler are scheduled to be executed prior to the main function being
executed. The init functions are great for setting up packages, initializing variables, or performing any other
bootstrapping you may need prior to the program running.

```go

```

## Tooling

```shell
go

Go is a tool for managing Go source code.

Usage:

        go <command> [arguments]

The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        work        workspace maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

        buildconstraint build constraints
        buildmode       build modes
        c               calling between Go and C
        cache           build and test caching
        environment     environment variables
        filetype        file types
        go.mod          the go.mod file
        gopath          GOPATH environment variable
        goproxy         module proxy protocol
        importpath      import path syntax
        modules         modules, module versions, and more
        module-auth     module authentication using go.sum
        packages        package lists and patterns
        private         configuration for downloading non-public code
        testflag        testing flags
        testfunc        testing functions
        vcs             controlling version control with GOVCS

```

### Variables declared without an initial value are set to their zero values:

* 0 for all integer types,
* 0.0 for floating point numbers,
* false for booleans,
* "" for strings,
* nil for interfaces, slices, channels, maps, pointers and functions.

Declare

```go
package main

import "fmt"

// The provided code is written in Go and demonstrates the concept of zero values
// in the language. Zero values are default values assigned to variables that are
// declared without an explicit initial value.
func main() {
	var n int
	var f float32
	var b bool
	var s string
	var sl []string

	fmt.Printf("n - type: %T value: %#v\n", n, n)
	fmt.Printf("f - type: %T value: %#v\n", f, f)
	fmt.Printf("b - type: %T value: %#v\n", b, b)
	fmt.Printf("s - type: %T value: %#v\n", s, s)
	fmt.Printf("sl - type: %T value: %#v\n", sl, sl)
}
```

## User-Define Types

Go allows you the ability to declare your own types using keyword struct.

Declaration of a struct type

```go
package main

import "fmt"

// user define-type in the program
type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

func main() {

	// Initialize a variable of type user by positional parameter
	// Must assign values to all fields.
	// Not prefer way because not knowing what the values are assigned to the fields
	rossi := user{"rossi", "rossi@email.com", 56, true}

	// Declare a variable of type user and initialize some or all the fields
	// prefer way because of code readability and knowing what the values are assigned to the fields
	lisa := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}

	// You may omit assign some values to the fields.
	// Omitted fields will be assigned the zero value of the field's type
	barry := user{
		name:  "barry",
		email: "barr@email.com",
	}

	fmt.Printf("%v, %T\n", rossi, rossi)
	fmt.Printf("%v, %T\n", lisa, lisa)
	fmt.Printf("%v, %T\n", barry, barry)
}
```

## Methods

* has no return value
* has receiver parameter declared between keyword word and the function name
* provide a way to add behavior to user-defined types
* provide a way to access interfaces

```go
// user defines a user in the program.
type user struct {
	name string
	email string
}

// notify implements a method with a value receiver.
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail implements a method with a pointer receiver.
func (u *user) changeEmail(email string) {
	u.email = email
}
```

## Function

## Access Modifier

In GO, there are two modifier types—package & public. An identifier that starts with a capital letter is public and
can be accessed by anyone outside the package. Whereas, an identifier that starts with a lowercase letter is package
access and can not be accessed from other packages.

Public access modifier

```go
package somepackage

// public access
func GetById(int id) int {}
var MyId
```

Package access modifier

```go
// package access - only methods and functions in this package can access them
func getById(int id) int {}
var myId
```

