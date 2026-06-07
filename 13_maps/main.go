// Maps store key-value pairs — like a dictionary or hash table
package main

import "fmt"

func main() {
	// Map literal
	capitals := map[string]string{
		"India":   "New Delhi",
		"Japan":   "Tokyo",
		"Germany": "Berlin",
	}

	fmt.Println("Capital of Japan:", capitals["Japan"])

	// Add or update a key
	capitals["Brazil"] = "Brasília"

	// Delete a key
	delete(capitals, "Germany")

	// Check if a key exists — "ok" is true if found
	city, ok := capitals["Germany"]
	if ok {
		fmt.Println("Germany:", city)
	} else {
		fmt.Println("Germany was deleted")
	}

	// Iterate over a map (order is random in Go)
	fmt.Println("--- All Capitals ---")
	for country, capital := range capitals {
		fmt.Printf("%s -> %s\n", country, capital)
	}

	// make creates an empty map
	scores := make(map[string]int)
	scores["Alice"] = 95
	scores["Bob"] = 80
	fmt.Println("Scores:", scores)
}
