// Test files end in _test.go — test functions start with Test and take *testing.T
package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("Add(2,3) = %d; want %d", got, want)
	}
}

// Table-driven tests are idiomatic Go — test many cases cleanly
func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{1, "1"},
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
		{30, "FizzBuzz"},
	}

	for _, tc := range tests {
		got := FizzBuzz(tc.input)
		if got != tc.want {
			t.Errorf("FizzBuzz(%d) = %q; want %q", tc.input, got, tc.want)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"racecar", true},
		{"hello", false},
		{"madam", true},
		{"a", true},
		{"", true},
	}

	for _, tc := range tests {
		got := IsPalindrome(tc.input)
		if got != tc.want {
			t.Errorf("IsPalindrome(%q) = %v; want %v", tc.input, got, tc.want)
		}
	}
}
