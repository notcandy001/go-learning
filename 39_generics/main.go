// Generics (Go 1.18+) let you write functions and types that work with multiple types
package main

import "fmt"

// [T comparable] means T can be any type that supports == and !=
func contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Number is a type constraint — only numeric types allowed
type Number interface {
	int | int32 | int64 | float32 | float64
}

func sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

// Generic Stack data structure
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

func (s *Stack[T]) Size() int { return len(s.items) }

func main() {
	ints := []int{1, 2, 3, 4, 5}
	strs := []string{"go", "rust", "zig"}

	fmt.Println("Contains 3:", contains(ints, 3))
	fmt.Println("Contains 9:", contains(ints, 9))
	fmt.Println("Contains 'go':", contains(strs, "go"))

	fmt.Println("Int sum:", sum(ints))
	fmt.Println("Float sum:", sum([]float64{1.1, 2.2, 3.3}))

	var stack Stack[string]
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")
	fmt.Println("Stack size:", stack.Size())
	if v, ok := stack.Pop(); ok {
		fmt.Println("Popped:", v)
	}
}
