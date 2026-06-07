// Methods are functions attached to a type — defined with a receiver between "func" and the name
package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

// Value receiver — gets a copy of the struct (fine for read-only methods)
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Pointer receiver — needed when you want to modify the struct
func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// String() method makes a type print nicely with fmt.Println
func (c Circle) String() string {
	return fmt.Sprintf("Circle(r=%.2f)", c.Radius)
}

func main() {
	c := Circle{Radius: 5}
	fmt.Println(c)
	fmt.Printf("Area: %.4f\n", c.Area())
	fmt.Printf("Perimeter: %.4f\n", c.Perimeter())

	c.Scale(2)
	fmt.Println("After Scale(2):", c)

	r := Rectangle{Width: 8, Height: 3}
	fmt.Printf("Rectangle Area: %.1f\n", r.Area())
	fmt.Printf("Rectangle Perimeter: %.1f\n", r.Perimeter())
}
