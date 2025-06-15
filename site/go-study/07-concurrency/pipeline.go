package main

import "fmt"

func main() {
	// input
	nums := []int{2, 3, 4, 7, 1}
	// stage 1
	dataChannel := sliceToChannel(nums)
	// stage 2
	finalChannel := square(dataChannel)
	//stage 3
	for n := range finalChannel {
		fmt.Printf("result: %v\n", n)
	}
}

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			fmt.Printf("input: %v\n", n)
			out <- n
		}
		close(out)
	}()

	return out
}

// square
// input: channel
// output: channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}
