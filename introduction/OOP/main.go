package main

import (
	"fmt"
	"math"
)

// Define the Shape interface with a method Area
type Shape interface {
	Area() float64
}

// Define a Circle struct
type Circle struct {
	Radius float64
}

// Define a Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Method for Circle to satisfy the Shape interface
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Method for Rectangle to satisfy the Shape interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// A function that accepts any type that satisfies the Shape interface
func PrintArea(s Shape) {
	fmt.Printf("The area of the shape is: %.2f\n", s.Area())
}

func main() {
	// Create instances of Circle and Rectangle
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 5}

	// Print the area of both shapes using the same function
	PrintArea(circle)    // Circle
	PrintArea(rectangle) // Rectangle
}
