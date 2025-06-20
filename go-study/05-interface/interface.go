package main

import "fmt"

func main() {
	var w Writer = ConsoleWriter{}
	_, error := w.Write([]byte("Hello GO interface"))

	if error != nil {
		fmt.Println("Error occured")
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

// The ConsoleWriter struct is defined without any properties. It implements the
// Write method declared by the Writer interface. The Write method for
// ConsoleWriter converts the byte slice to a string and prints it to the console
// using fmt.Println. It then returns the number of bytes written and any error
// that occurred.
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
