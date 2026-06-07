// Interfaces define a set of method signatures — any type that implements them satisfies the interface
package main

import (
	"fmt"
	"math"
)

// The Shape interface — any type with Area() and Perimeter() is a Shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct{ Radius float64 }
type Rectangle struct{ Width, Height float64 }
type Triangle struct{ A, B, C float64 }

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

// Heron's formula for triangle area
func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}
func (t Triangle) Perimeter() float64 { return t.A + t.B + t.C }

// printShapeInfo accepts any Shape — this is polymorphism in Go
func printShapeInfo(s Shape) {
	fmt.Printf("Type: %T | Area: %.2f | Perimeter: %.2f\n", s, s.Area(), s.Perimeter())
}

func main() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
		Triangle{A: 3, B: 4, C: 5},
	}

	for _, s := range shapes {
		printShapeInfo(s)
	}
}
