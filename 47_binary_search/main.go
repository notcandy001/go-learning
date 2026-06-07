// Binary search finds a value in a sorted slice in O(log n) — much faster than linear search
package main

import "fmt"

// BinarySearch returns the index of target, or -1 if not found
// The slice must be sorted in ascending order
func BinarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2 // avoids integer overflow

		if nums[mid] == target {
			return mid // found it!
		} else if nums[mid] < target {
			low = mid + 1 // target is in the right half
		} else {
			high = mid - 1 // target is in the left half
		}
	}
	return -1 // not found
}

// Recursive version
func BinarySearchRecursive(nums []int, target, low, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return BinarySearchRecursive(nums, target, mid+1, high)
	}
	return BinarySearchRecursive(nums, target, low, mid-1)
}

func main() {
	sorted := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	fmt.Println("Slice:", sorted)

	targets := []int{23, 56, 1, 91, 37}
	for _, t := range targets {
		idx := BinarySearch(sorted, t)
		if idx != -1 {
			fmt.Printf("Found %d at index %d\n", t, idx)
		} else {
			fmt.Printf("%d not found\n", t)
		}
	}

	fmt.Println("\nRecursive:")
	fmt.Println("72:", BinarySearchRecursive(sorted, 72, 0, len(sorted)-1))
}
