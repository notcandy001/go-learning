// The sort package provides sorting for slices and custom types
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface for []Person based on Age
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	// Sort ints
	nums := []int{5, 2, 9, 1, 7, 3}
	sort.Ints(nums)
	fmt.Println("Sorted ints:", nums)

	// Sort strings
	words := []string{"banana", "apple", "cherry", "date"}
	sort.Strings(words)
	fmt.Println("Sorted strings:", words)

	// Sort in reverse using sort.Reverse
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println("Reverse sorted:", nums)

	// Custom sort with sort.Slice — simplest approach for structs
	people := []Person{
		{"Candy", 19}, {"Alice", 25}, {"Bob", 17}, {"Zara", 22},
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("By age:", people)

	// Sort by name
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Println("By name:", people)

	// sort.Interface approach (more verbose but satisfies the interface)
	sort.Sort(ByAge(people))
	fmt.Println("By age (interface):", people)

	// Check if sorted
	fmt.Println("Is sorted:", sort.IntsAreSorted([]int{1, 2, 3, 4}))
}
