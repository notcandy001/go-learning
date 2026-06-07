// net/http lets Go make HTTP requests — GET, POST, and more
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Struct matching a public API response
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// Simple GET request
	url := "https://jsonplaceholder.typicode.com/todos/1"

	// Use a client with a timeout — never use http.Get without one in real code
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer resp.Body.Close() // always close the body

	fmt.Println("Status:", resp.Status)
	fmt.Println("Content-Type:", resp.Header.Get("Content-Type"))

	// Read the raw body
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Raw body:", string(body))

	// Decode JSON response into struct
	resp2, _ := client.Get(url)
	defer resp2.Body.Close()
	var todo Todo
	json.NewDecoder(resp2.Body).Decode(&todo)
	fmt.Printf("Todo: [%d] %s (done=%v)\n", todo.ID, todo.Title, todo.Completed)
}
