// Constants are values that never change — declared with the "const" keyword
package main

import "fmt"

// Constants can be declared at the package level (outside any function)
const Pi = 3.14159
const AppName = "GoLearner"
const MaxRetries = 5

func main() {
	// You can also declare a block of constants together
	const (
		StatusOK    = 200
		StatusNotFound = 404
	)

	fmt.Println("App:", AppName)
	fmt.Println("Pi:", Pi)
	fmt.Println("Max Retries:", MaxRetries)
	fmt.Println("HTTP 200:", StatusOK)
	fmt.Println("HTTP 404:", StatusNotFound)
}
