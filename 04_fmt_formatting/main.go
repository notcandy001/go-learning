// The fmt package has powerful verbs for formatting output — like printf in C
package main

import "fmt"

func main() {
	name := "Candy"
	age := 19
	gpa := 9.1
	isHappy := true

	// %s = string, %d = integer, %f = float, %t = boolean, %v = any value
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("GPA: %.2f\n", gpa) // .2 means 2 decimal places
	fmt.Printf("Happy: %t\n", isHappy)

	// Sprintf returns a formatted string instead of printing it
	greeting := fmt.Sprintf("Hello, %s! You are %d years old.", name, age)
	fmt.Println(greeting)

	// %T prints the type of a variable
	fmt.Printf("Type of age: %T\n", age)
	fmt.Printf("Type of gpa: %T\n", gpa)

	// %v is the "default" format — works for any type
	fmt.Printf("All values: %v, %v, %v\n", name, age, gpa)
}
