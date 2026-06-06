// Go has only one loop keyword — "for" — but it can work like for, while, and foreach
package main

import "fmt"

func main() {
	// Classic C-style for loop: init; condition; post
	fmt.Println("--- Counting 1 to 5 ---")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Loop over a slice with range — gives index and value
	fmt.Println("--- Fruits ---")
	fruits := []string{"mango", "papaya", "lychee"}
	for index, fruit := range fruits {
		fmt.Printf("[%d] %s\n", index, fruit)
	}

	// Ignore the index with _ (blank identifier)
	fmt.Println("--- Just values ---")
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// Loop over a string — range gives rune (character) values
	fmt.Println("--- Letters in 'Go' ---")
	for i, ch := range "Go" {
		fmt.Printf("index %d: %c\n", i, ch)
	}

	// Loop over a map
	fmt.Println("--- Map ---")
	scores := map[string]int{"Alice": 90, "Bob": 85}
	for name, score := range scores {
		fmt.Printf("%s scored %d\n", name, score)
	}
}
