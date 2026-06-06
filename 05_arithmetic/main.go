// Go supports all basic math operators — +, -, *, /, and % (modulo/remainder)
package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 17, 5

	// Basic operators
	fmt.Println("Add:      ", a+b)
	fmt.Println("Subtract: ", a-b)
	fmt.Println("Multiply: ", a*b)
	fmt.Println("Divide:   ", a/b)   // integer division — remainder is dropped
	fmt.Println("Modulo:   ", a%b)   // remainder after division

	// For true division, use float64
	fmt.Printf("True div: %.2f\n", float64(a)/float64(b))

	// Increment and decrement
	x := 10
	x++
	fmt.Println("After x++:", x)
	x--
	fmt.Println("After x--:", x)

	// math package for more operations
	fmt.Println("Square root of 144:", math.Sqrt(144))
	fmt.Println("2 to the power 8:  ", math.Pow(2, 8))
	fmt.Println("Absolute value -7: ", math.Abs(-7))
}
