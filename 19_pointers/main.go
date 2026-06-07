// Pointers store the memory address of a variable — & gets the address, * dereferences it
package main

import "fmt"

// Without pointer: changes don't affect the original
func doubleVal(n int) {
	n = n * 2
}

// With pointer: changes affect the original variable
func doublePtr(n *int) {
	*n = *n * 2 // * in front dereferences — means "the value at this address"
}

func main() {
	x := 10
	fmt.Println("x before:", x)

	doubleVal(x)
	fmt.Println("x after doubleVal:", x) // still 10 — copy was modified

	doublePtr(&x) // & means "address of x"
	fmt.Println("x after doublePtr:", x) // now 20

	// Pointer basics
	y := 42
	p := &y        // p is a pointer to y
	fmt.Println("Value of y:", y)
	fmt.Println("Address in p:", p)
	fmt.Println("Value at p:", *p) // dereference

	*p = 100 // change y through the pointer
	fmt.Println("y after *p = 100:", y)

	// new() allocates a zero-value and returns a pointer
	ptr := new(int)
	*ptr = 55
	fmt.Println("new(int):", *ptr)
}
