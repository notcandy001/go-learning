// Functions are reusable blocks of code — defined with the "func" keyword
package main

import "fmt"

// A simple function that takes no parameters and returns nothing
func greet() {
	fmt.Println("Hello from a function!")
}

// A function with parameters and a return value
func add(a int, b int) int {
	return a + b
}

// If parameters share a type, you can write it once
func multiply(a, b int) int {
	return a * b
}

// Functions are first-class — you can pass them as arguments
func applyTwice(f func(int) int, x int) int {
	return f(f(x))
}

func double(n int) int {
	return n * 2
}

func main() {
	greet()

	sum := add(7, 3)
	fmt.Println("7 + 3 =", sum)
	fmt.Println("6 * 9 =", multiply(6, 9))

	// Pass a function as an argument
	result := applyTwice(double, 5)
	fmt.Println("double(double(5)) =", result)

	// Anonymous function (defined and called immediately)
	square := func(n int) int {
		return n * n
	}
	fmt.Println("Square of 8:", square(8))
}
