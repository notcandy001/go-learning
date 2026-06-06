// switch is a cleaner alternative to long if/else chains — no "break" needed in Go
package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Now().Weekday().String()

	// Basic switch on a string value
	switch day {
	case "Monday":
		fmt.Println("Start of the week — let's go!")
	case "Friday":
		fmt.Println("Almost the weekend!")
	case "Saturday", "Sunday":
		// Multiple values in one case
		fmt.Println("Weekend! Rest up.")
	default:
		fmt.Println("Midweek grind:", day)
	}

	// Switch with no expression — acts like if/else
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Excellent")
	case score >= 70:
		fmt.Println("Good")
	default:
		fmt.Println("Needs improvement")
	}

	// Type switch — checks the underlying type of an interface value
	var val interface{} = 42
	switch v := val.(type) {
	case int:
		fmt.Println("It's an int:", v)
	case string:
		fmt.Println("It's a string:", v)
	default:
		fmt.Println("Unknown type")
	}
}
