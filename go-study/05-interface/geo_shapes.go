// This code demonstrates the power of interfaces in Go, allowing for polymorphic
// behavior where the specific type of the object isn't important, only that it
// adheres to a certain interface. In this case, any shape that can calculate its
// area can be used with the getArea function.
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

// Area
// Both Rectangle and Circle structs implement the Area method declared by
// the Shape interface. The Area method for Rectangle returns the product of its
// Width and Height, while for Circle, it returns the product of Pi, Radius and
// Radius.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func getArea(shape Shape) {
	fmt.Println(shape.Area())
}

func main() {
	r := Rectangle{Width: 7, Height: 8}
	c := Circle{Radius: 5}

	getArea(r)
	getArea(c)
}
