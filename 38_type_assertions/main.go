// Type assertions extract the concrete type from an interface value
package main

import "fmt"

func describe(i interface{}) {
	// Two-value assertion — safe, doesn't panic
	if s, ok := i.(string); ok {
		fmt.Printf("String of length %d: %q\n", len(s), s)
		return
	}
	if n, ok := i.(int); ok {
		fmt.Printf("Integer: %d (double = %d)\n", n, n*2)
		return
	}
	fmt.Printf("Unknown type: %T = %v\n", i, i)
}

type Stringer interface {
	String() string
}

type Point struct{ X, Y int }

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func printIfStringer(val interface{}) {
	// Check if val implements the Stringer interface
	if s, ok := val.(Stringer); ok {
		fmt.Println("Stringer:", s.String())
	} else {
		fmt.Println("Not a Stringer:", val)
	}
}

func main() {
	var values []interface{} = []interface{}{
		"hello", 42, 3.14, true, []int{1, 2, 3},
	}

	for _, v := range values {
		describe(v)
	}

	p := Point{3, 4}
	printIfStringer(p)
	printIfStringer("just a string")

	// One-value assertion — panics if wrong type, so only use when you're certain
	var i interface{} = "definitely a string"
	s := i.(string)
	fmt.Println("Asserted:", s)
}
