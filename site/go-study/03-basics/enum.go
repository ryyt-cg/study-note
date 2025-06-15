package main

import "fmt"

type Direction int

const (
	North Direction = iota
	NorthEast
	East
	SouthEast
	South
	SouthWest
	West
	NorthWest
)

func (d Direction) Ok() bool {
	switch d {
	case North,
		NorthEast,
		East,
		SouthEast,
		South,
		SouthWest,
		West,
		NorthWest:
		return true
	}
	return false
}

func (d Direction) String() string {
	return [...]string{"North", "North East", "South East", "East", "South", "South West", "West", "North West"}[d]
}

func main() {
	for dir := Direction(0); dir.Ok(); dir++ {
		fmt.Printf("%d: %s\n", dir, dir.String())
	}
}
