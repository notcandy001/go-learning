// Variadic functions accept any number of arguments using "..." before the type
package main

import "fmt"

// ...int means you can pass zero or more ints — they arrive as a slice
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Mix normal and variadic params — variadic must be last
func logMessage(level string, parts ...string) {
	fmt.Printf("[%s]", level)
	for _, p := range parts {
		fmt.Print(" " + p)
	}
	fmt.Println()
}

func main() {
	fmt.Println("sum(1,2,3):", sum(1, 2, 3))
	fmt.Println("sum(1..10):", sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println("sum():", sum()) // zero arguments — valid!

	// You can also spread a slice into a variadic function with ...
	numbers := []int{10, 20, 30}
	fmt.Println("sum(slice...):", sum(numbers...))

	logMessage("INFO", "Server", "started", "on port 8080")
	logMessage("ERROR", "File not found")
}
