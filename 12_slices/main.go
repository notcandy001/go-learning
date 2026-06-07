// Slices are dynamic arrays — the most commonly used collection in Go
package main

import "fmt"

func main() {
	// Create a slice with a literal
	nums := []int{10, 20, 30, 40, 50}
	fmt.Println("Slice:", nums)
	fmt.Println("Len:", len(nums), "Cap:", cap(nums))

	// append adds elements — Go grows the underlying array automatically
	nums = append(nums, 60, 70)
	fmt.Println("After append:", nums)

	// Slicing a slice: [low:high] — low is inclusive, high is exclusive
	part := nums[1:4]
	fmt.Println("nums[1:4]:", part)

	// make creates a slice with a given length and optional capacity
	s := make([]int, 3, 5)
	fmt.Println("make slice:", s, "len:", len(s), "cap:", cap(s))

	// 2D slice (slice of slices)
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix row 1:", matrix[1])

	// Copy a slice — unlike arrays, slices share memory if you use =
	original := []int{1, 2, 3}
	clone := make([]int, len(original))
	copy(clone, original) // safe independent copy
	clone[0] = 99
	fmt.Println("original:", original)
	fmt.Println("clone:", clone)
}
