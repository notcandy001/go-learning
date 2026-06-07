// A closure is a function that "closes over" (remembers) variables from its surrounding scope
package main

import "fmt"

// makeCounter returns a function that increments an internal counter each call
func makeCounter() func() int {
	count := 0 // this variable lives inside the closure
	return func() int {
		count++
		return count
	}
}

// makeAdder creates a function that always adds a fixed number
func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y // x is "closed over" from makeAdder's scope
	}
}

func main() {
	// Each counter has its own independent "count" variable
	counterA := makeCounter()
	counterB := makeCounter()

	fmt.Println(counterA()) // 1
	fmt.Println(counterA()) // 2
	fmt.Println(counterA()) // 3
	fmt.Println(counterB()) // 1 — B is independent of A

	add5 := makeAdder(5)
	add10 := makeAdder(10)
	fmt.Println("add5(3):", add5(3))   // 8
	fmt.Println("add10(3):", add10(3)) // 13

	// Closure capturing a loop variable — common gotcha
	// Always capture by value if you use closures inside loops
	fns := make([]func(), 3)
	for i := 0; i < 3; i++ {
		i := i // shadow i — creates a new variable per iteration
		fns[i] = func() { fmt.Println("i =", i) }
	}
	for _, f := range fns {
		f()
	}
}
