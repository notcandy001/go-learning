// This file contains logic to test — see math_test.go for the actual tests
// Run tests with: go test ./...
package main

import "fmt"

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// FizzBuzz returns the FizzBuzz string for a number
func FizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", n)
	}
}

// IsPalindrome checks if a string reads the same forwards and backwards
func IsPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Run: go test -v to see test results")
	fmt.Println("Add(2,3):", Add(2, 3))
	fmt.Println("FizzBuzz(15):", FizzBuzz(15))
	fmt.Println("IsPalindrome('racecar'):", IsPalindrome("racecar"))
}
