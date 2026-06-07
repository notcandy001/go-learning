// The blank identifier _ discards values you don't need — Go requires all declared variables to be used
package main

import (
	"fmt"
	"os"
)

func threeValues() (int, string, bool) {
	return 42, "hello", true
}

func main() {
	// Discard the second return value — without _ this would be a compile error
	n, _, flag := threeValues()
	fmt.Println("n:", n, "flag:", flag)

	// Discard the index when ranging
	fruits := []string{"mango", "lychee", "papaya"}
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// Discard error (not recommended in real code, but legal)
	_ = os.Setenv("DEMO", "value")

	// Import for side effects only — runs the package's init() but doesn't use exports
	// import _ "image/png"  // registers PNG decoder without using the package directly

	// Discard map value when you only need to check existence
	m := map[string]int{"a": 1, "b": 2}
	_, exists := m["c"]
	fmt.Println("'c' exists:", exists)

	// Blank identifier can also suppress "declared but not used" errors during dev
	_ = "work in progress"
}
