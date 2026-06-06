// if/else lets your program make decisions based on conditions
package main

import "fmt"

func main() {
	score := 78

	// Basic if/else if/else chain
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else if score >= 60 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// Go allows a short statement before the condition (scoped to the if block)
	if remainder := score % 10; remainder > 5 {
		fmt.Println("Rounds up to next grade boundary")
	} else {
		fmt.Println("Rounds down, remainder:", remainder)
	}

	// Checking multiple conditions with && (AND) and || (OR)
	age := 20
	hasID := true
	if age >= 18 && hasID {
		fmt.Println("Access granted")
	}

	// Negation with !
	isRaining := false
	if !isRaining {
		fmt.Println("No umbrella needed")
	}
}
