// The os and bufio packages let you read and write files on disk
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "/tmp/go_learning_test.txt"

	// Write a file — os.WriteFile is the simplest way (Go 1.16+)
	content := "Hello from Go!\nLine 2\nLine 3\n"
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}
	fmt.Println("File written:", filename)

	// Read the entire file at once
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	fmt.Println("Full content:\n" + string(data))

	// Read line by line using bufio.Scanner — good for large files
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open error:", err)
		return
	}
	defer file.Close() // always close files — defer makes sure it happens

	fmt.Println("--- Line by line ---")
	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fmt.Printf("%d: %s\n", lineNum, line)
		lineNum++
	}

	// Append to a file — os.O_APPEND flag
	f, _ := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	f.WriteString("Appended line\n")
	fmt.Println("Line appended")

	// Clean up
	os.Remove(filename)
}
