// Variables hold data — Go is statically typed, so each variable has a fixed type
package main

import "fmt"

func main() {
	// Long form: declare type explicitly
	var name string = "Candy"

	// Short form: Go infers the type automatically (:= only works inside functions)
	age := 19

	// Declare without value — Go sets a "zero value" (0, "", false, etc.)
	var score int

	// Booleans are true or false
	var isLearning bool = true

	// float64 for decimal numbers
	var gpa float64 = 9.1

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Score (zero value):", score)
	fmt.Println("Is Learning:", isLearning)
	fmt.Println("GPA:", gpa)
}
