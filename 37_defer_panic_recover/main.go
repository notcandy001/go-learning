// defer runs a function just before the enclosing function returns — great for cleanup
// panic stops normal execution; recover catches a panic in a deferred function
package main

import "fmt"

func cleanupDemo() {
	defer fmt.Println("3. Cleanup runs last — even though it was registered first")
	defer fmt.Println("2. Second defer (LIFO order)")
	fmt.Println("1. Function body runs")
}

func safeDiv(a, b int) (result int, err error) {
	// recover must be called inside a deferred function
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			err = fmt.Errorf("caught panic: %v", r)
		}
	}()

	if b == 0 {
		panic("attempted division by zero!") // triggers panic
	}
	return a / b, nil
}

func riskyOperation() {
	fmt.Println("Before panic")
	panic("something went terribly wrong")
	fmt.Println("This line never runs")
}

func main() {
	cleanupDemo()
	fmt.Println()

	result, err := safeDiv(10, 2)
	fmt.Println("10/2 =", result, "err:", err)

	result2, err2 := safeDiv(10, 0)
	fmt.Println("10/0 =", result2, "err:", err2)

	// Demonstrate recover in a goroutine wrapper
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Main recovered:", r)
		}
	}()
	riskyOperation()
}
