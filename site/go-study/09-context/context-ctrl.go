package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Contexts are widely used in various parts of Go applications,
// including network servers, databases, and client requests.
// hey help properly manage task execution time, cancel unnecessary operations,
// and pass data between goroutines.
func main() {
	timeoutContext, cancel := context.WithTimeout(context.Background(), 900*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutContext, "GET", "https://placehold.it/2000x2000", nil)
	if err != nil {
		panic(err)
	}

	// Create a new context with the request context as the parent.
	//ctx := req.Context()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	imageData, err := ioutil.ReadAll(res.Body)
	fmt.Printf("download image of size %d\n", len(imageData))
}
