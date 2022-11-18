package main

import (
	"fmt"
)

// Creating an Interface
type Shape interface {
	// Method Signatures
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

// Implementing Methods of the Shape Interface
func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

// Main Method
func main() {
	// Accessing Elements of the Shape Interface
	var s Shape
	s = Rectangle{10, 14}
	fmt.Println("Area of Shape :", s.Area())
	fmt.Println("Perimeter of Shape:", s.Perimeter())
}

/*
OUTPUT
Area of Shape : 140
Perimeter of Shape: 48
*/
