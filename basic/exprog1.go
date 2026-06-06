package main

// this is to declare the main function

import "fmt"

// this is to declare the import func

func main() {
	// this is to declare the main function

	var name string
	var age int
	var homeaddress string
	var Phonenumber int

	// this is to declare the ask the user input
	fmt.Println("enter ur name ")
	fmt.Scanln(&name)

	fmt.Println("enter ur age")
	fmt.Scanln(&age)

	fmt.Println("enter ur home address")
	fmt.Scanln(&homeaddress)

	fmt.Println("enter ur phone number")
	fmt.Scanln(&Phonenumber)

	// this is to declare the output func

	fmt.Println("hello", name)
	fmt.Println("ur age is ", age)
	fmt.Println("ur from", homeaddress)
	fmt.Println("ur number is ", Phonenumber)

}
