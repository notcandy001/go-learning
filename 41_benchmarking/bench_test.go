// Benchmark functions start with Benchmark and take *testing.B
package main

import "testing"

func BenchmarkConcatPlus(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N is determined automatically by the test runner
		concatPlus(100)
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatBuilder(100)
	}
}
