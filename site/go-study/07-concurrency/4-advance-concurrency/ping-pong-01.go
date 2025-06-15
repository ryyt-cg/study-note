package main

import "time"

type Ball struct {
	hits int
}

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball) // game on; toss the ball, deadlock if this line comments out
	time.Sleep(1 * time.Second)
	<-table // game over; grab the ball

	panic("show me the stacks")
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
