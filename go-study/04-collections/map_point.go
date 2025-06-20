package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	m := make(map[Point]int)

	m[Point{1, 2}] = 1
	fmt.Println(m[Point{1, 2}]) // 1
}
