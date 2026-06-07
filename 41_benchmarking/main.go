// Go's testing package includes benchmarking — run with: go test -bench=. -benchmem
package main

import "fmt"

// Two ways to concatenate strings — one is much faster
func concatPlus(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "x" // slow — creates a new string each iteration
	}
	return s
}

func concatBuilder(n int) string {
	// strings.Builder pre-allocates — much faster for many concatenations
	var b []byte
	for i := 0; i < n; i++ {
		b = append(b, 'x')
	}
	return string(b)
}

func main() {
	fmt.Println("Run: go test -bench=. -benchmem to compare performance")
	fmt.Println("Plus len:", len(concatPlus(100)))
	fmt.Println("Builder len:", len(concatBuilder(100)))
}
