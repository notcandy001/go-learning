// net/http also lets you build HTTP servers — Go's standard library is surprisingly powerful here
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// Handler functions receive (ResponseWriter, *Request)
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Go HTTP server!")
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // get query param ?name=...
	if name == "" {
		name = "stranger"
	}
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := Response{Message: "OK", Status: 200}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// Register routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/greet", greetHandler) // try: curl "localhost:8080/greet?name=Candy"
	http.HandleFunc("/json", jsonHandler)   // try: curl localhost:8080/json

	fmt.Println("Server running at http://localhost:8080")
	fmt.Println("Try: curl localhost:8080/greet?name=Candy")

	// ListenAndServe blocks — Ctrl+C to stop
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
