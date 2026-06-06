// Go has no "while" keyword — a for loop with only a condition works exactly like while
package main

import "fmt"

func main() {
	// while-style loop: keep going until condition is false
	count := 0
	for count < 5 {
		fmt.Println("count:", count)
		count++
	}

	// Infinite loop with break — useful for game loops, servers, REPL shells
	n := 1
	fmt.Println("--- Powers of 2 below 100 ---")
	for {
		if n >= 100 {
			break // exit the loop
		}
		fmt.Println(n)
		n *= 2
	}

	// continue skips the rest of the current iteration and goes to the next
	fmt.Println("--- Even numbers 1-10 ---")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue // skip odd numbers
		}
		fmt.Println(i)
	}
}
