// Custom error types give you structured errors with extra fields you can inspect
package main

import (
	"errors"
	"fmt"
)

// Define a custom error type by implementing the error interface (just needs Error() string)
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field %q: %s", e.Field, e.Message)
}

type NotFoundError struct {
	Resource string
	ID       int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID %d not found", e.Resource, e.ID)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "must be non-negative"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: "unrealistically large"}
	}
	return nil
}

func findUser(id int) (string, error) {
	users := map[int]string{1: "Alice", 2: "Bob"}
	name, ok := users[id]
	if !ok {
		return "", &NotFoundError{Resource: "User", ID: id}
	}
	return name, nil
}

func main() {
	// errors.As extracts the concrete error type if it matches
	err := validateAge(-5)
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Printf("Field: %s | Msg: %s\n", ve.Field, ve.Message)
	}

	name, err2 := findUser(3)
	var nfe *NotFoundError
	if errors.As(err2, &nfe) {
		fmt.Printf("Not found: %s #%d\n", nfe.Resource, nfe.ID)
	} else {
		fmt.Println("Found:", name)
	}

	// Successful case
	user, err3 := findUser(1)
	if err3 == nil {
		fmt.Println("Found user:", user)
	}
}
