// Arrays have a fixed size set at compile time — rarely used directly, but good to understand
package main

import "fmt"

func main() {
	// Declare an array of 5 ints — all zero by default
	var grades [5]int
	grades[0] = 92
	grades[1] = 88
	grades[2] = 76
	grades[3] = 95
	grades[4] = 83

	fmt.Println("Grades:", grades)
	fmt.Println("Length:", len(grades))

	// Array literal — declare and initialize at once
	colors := [3]string{"red", "green", "blue"}
	fmt.Println("Colors:", colors)

	// Use ... to let Go count the elements
	primes := [...]int{2, 3, 5, 7, 11}
	fmt.Println("Primes:", primes)
	fmt.Println("Third prime:", primes[2])

	// Arrays are value types — copying creates an independent copy
	a := [3]int{1, 2, 3}
	b := a      // b is a separate copy
	b[0] = 99
	fmt.Println("a:", a) // still {1,2,3}
	fmt.Println("b:", b) // {99,2,3}
}
