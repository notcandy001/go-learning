package main

import "fmt"

func main() {

	var name string

	// this is to declare the  statement with loop

	fmt.Println("enter ur fav lanuage ")
	fmt.Scanln(&name)
	// this is to declare the for loop statement

	for i := 1; i <= 25; i++ {
		fmt.Println("my fav is ", name)

	}

}
