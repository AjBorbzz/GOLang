package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method with a pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func main() {
	rect := Rectangle{Width: 4, Height: 3}
	area := rect.Area()
	fmt.Printf("Area: %.2f\n", area)

	rect.Scale(2)
	fmt.Printf("New Width: %.2f, New Height: %.2f\n", rect.Width, rect.Height)
	fmt.Printf("New Area: %.2f\n", rect.Area())
}
