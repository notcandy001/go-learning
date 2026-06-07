// Go functions can return multiple values — this is idiomatic Go, especially for errors
package main

import (
	"errors"
	"fmt"
)

// Returns both the quotient and remainder
func divide(a, b int) (int, int) {
	return a / b, a % b
}

// Named return values — the return variables are declared in the signature
func minMax(nums []int) (min, max int) {
	min, max = nums[0], nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return // "naked return" uses the named variables
}

// The classic Go pattern: return (result, error)
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	q, r := divide(17, 5)
	fmt.Printf("17 / 5 = %d remainder %d\n", q, r)

	nums := []int{42, 7, 99, 3, 55}
	lo, hi := minMax(nums)
	fmt.Printf("Min: %d, Max: %d\n", lo, hi)

	result, err := safeDivide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.4f\n", result)
	}

	_, err2 := safeDivide(5, 0)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
}
