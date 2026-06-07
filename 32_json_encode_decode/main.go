// encoding/json marshals Go structs to JSON and unmarshals JSON back into structs
package main

import (
	"encoding/json"
	"fmt"
)

// JSON struct tags control how fields appear in JSON output
type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Email   string   `json:"email,omitempty"` // omitted if empty
	private string   // unexported fields are ignored by json package
	Tags    []string `json:"tags"`
}

func main() {
	// Marshal: Go struct → JSON bytes
	p := Person{
		Name:  "Candy",
		Age:   19,
		Email: "candy@example.com",
		Tags:  []string{"linux", "go", "rice"},
	}

	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}
	fmt.Println("JSON:", string(data))

	// MarshalIndent for pretty-printed JSON
	pretty, _ := json.MarshalIndent(p, "", "  ")
	fmt.Println("Pretty JSON:\n" + string(pretty))

	// Unmarshal: JSON bytes → Go struct
	jsonStr := `{"name":"Nisa","age":20,"tags":["qml","wayland"]}`
	var p2 Person
	err = json.Unmarshal([]byte(jsonStr), &p2)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}
	fmt.Printf("Decoded: %+v\n", p2)

	// Decode into a map when you don't know the structure ahead of time
	var raw map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &raw)
	fmt.Println("Name from map:", raw["name"])
	fmt.Println("Tags from map:", raw["tags"])
}
