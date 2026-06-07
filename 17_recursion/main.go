// Recursion is when a function calls itself — great for problems with repeated sub-structure
package main

import "fmt"

// Factorial: n! = n * (n-1) * ... * 1
func factorial(n int) int {
	if n <= 1 {
		return 1 // base case — stops the recursion
	}
	return n * factorial(n-1) // recursive case
}

// Fibonacci: each number is the sum of the two before it
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// Recursive sum of a slice
func sliceSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + sliceSum(nums[1:])
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d! = %d\n", i, factorial(i))
	}

	fmt.Println("--- Fibonacci sequence ---")
	for i := 0; i <= 10; i++ {
		fmt.Printf("fib(%d) = %d\n", i, fib(i))
	}

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Recursive sum:", sliceSum(nums))
}
