// Go doesn't use exceptions — functions return errors as values and you check them explicitly
package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Functions return (value, error) — caller must check the error
func parseInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parseInt failed for %q: %w", s, err)
	}
	return n, nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	// Happy path
	n, err := parseInt("42")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed:", n)
	}

	// Error path
	_, err = parseInt("abc")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// errors.Is checks if an error matches a target (works through wrapping)
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Divide error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// fmt.Errorf with %w wraps the original error
	_, err2 := parseInt("xyz")
	fmt.Println("Wrapped error:", err2)
	fmt.Println("Is strconv error?", errors.As(err2, &strconv.NumError{}))
}
