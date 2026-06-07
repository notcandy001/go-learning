// Goroutines are lightweight concurrent functions — start one with the "go" keyword
package main

import (
	"fmt"
	"time"
)

func printNumbers(label string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("[%s] %d\n", label, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func say(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg)
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	// Without "go" — runs sequentially, one after the other
	// say("hello")
	// say("world")

	// With "go" — both run concurrently
	go say("hello")
	go say("world")

	// IMPORTANT: main must wait or goroutines get killed when main exits
	// We'll use time.Sleep here; better solutions are in 29_waitgroup
	time.Sleep(500 * time.Millisecond)

	fmt.Println("--- Two concurrent counters ---")
	go printNumbers("A")
	go printNumbers("B")
	time.Sleep(700 * time.Millisecond)

	fmt.Println("Main done")
}
