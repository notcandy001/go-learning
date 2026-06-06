package main

import "fmt"

func main() {
	// this is to declare the main function
	var age int

	// this is to declare the user to declare the value

	fmt.Println("enter ur age")
	fmt.Scanln(&age)

	// this is to declare the loop func
	if age >= 20 {
		fmt.Println("the user is adult")
	} else {
		fmt.Println("user is kid")
	}

	// this is to declare the main function is done
	fmt.Println("done")

}
