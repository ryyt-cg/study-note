package main

import (
	"fmt"
)

type Animal interface {
	Sound() string
}

type Dog struct {
}

func (d Dog) Sound() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Sound() string {
	return "Meow!"
}

type Cow struct {
}

func (l Cow) Sound() string {
	return "Moo!"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Cow{}}

	for _, animal := range animals {
		fmt.Println(animal.Sound())
	}
}
