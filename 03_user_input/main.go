// fmt.Scan and bufio.Scanner let you read input typed by the user in the terminal
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Simple single-word input using fmt.Scan
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name) // & means "store into this variable"

	// bufio.Scanner reads a full line including spaces
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your favourite language: ")
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	fmt.Printf("Hey %s! So you like %s — great choice.\n", name, line)
}
