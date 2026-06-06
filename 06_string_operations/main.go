// Strings in Go are immutable sequences of bytes — the strings package helps manipulate them
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "  Hello, Gopher!  "

	// Length (in bytes)
	fmt.Println("Length:", len(strings.TrimSpace(s)))

	// Case conversion
	fmt.Println("Upper:", strings.ToUpper(s))
	fmt.Println("Lower:", strings.ToLower(s))

	// Trim whitespace from both ends
	fmt.Println("Trimmed:", strings.TrimSpace(s))

	// Check if a string contains a substring
	fmt.Println("Contains 'Gopher':", strings.Contains(s, "Gopher"))

	// Replace parts of a string
	fmt.Println("Replace:", strings.Replace(s, "Gopher", "World", 1))

	// Split a string into a slice
	csv := "apple,banana,cherry"
	fruits := strings.Split(csv, ",")
	fmt.Println("Fruits:", fruits)

	// Join a slice back into a string
	joined := strings.Join(fruits, " | ")
	fmt.Println("Joined:", joined)

	// String concatenation with +
	first := "Go"
	second := "lang"
	fmt.Println("Concat:", first+second)

	// Index access — returns a byte value
	fmt.Printf("First char byte: %c\n", s[2])
}
