// Structs group related fields together — Go's way of creating custom types (like classes without methods)
package main

import "fmt"

// Define a struct type — by convention, exported types start with a capital letter
type Person struct {
	Name string
	Age  int
	City string
}

// Nested struct
type Address struct {
	Street string
	Zip    string
}

type Employee struct {
	Person          // embedded — Employee "inherits" Person's fields
	Company string
	Address Address // nested struct as a field
}

func main() {
	// Struct literal — named fields (preferred, order doesn't matter)
	p := Person{Name: "Candy", Age: 19, City: "Mysuru"}
	fmt.Println("Person:", p)
	fmt.Println("Name:", p.Name)

	// Positional literal — order must match struct definition exactly
	p2 := Person{"Nisa", 20, "Tokyo"}
	fmt.Println("Person2:", p2)

	// Zero value struct — all fields are zero/empty
	var empty Person
	fmt.Println("Empty:", empty)

	// Pointer to struct — & creates a pointer, . still works for field access
	pp := &Person{Name: "Anon", Age: 25}
	pp.Age = 26 // Go auto-dereferences — no need to write (*pp).Age
	fmt.Println("Pointer person age:", pp.Age)

	// Nested struct access
	emp := Employee{
		Person:  Person{"Dev", 22, "Bangalore"},
		Company: "Brainitech",
		Address: Address{"MG Road", "570001"},
	}
	fmt.Println("Employee:", emp.Name, "at", emp.Company, "—", emp.Address.Street)
}
