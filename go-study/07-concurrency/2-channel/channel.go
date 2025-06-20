package main

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello"
		println("Sent!")
	}()

	// block until a message is received
	msg := <-ch
	println(msg)
}
