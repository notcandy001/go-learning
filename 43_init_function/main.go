// The init() function runs automatically before main() — useful for setup and registration
package main

import "fmt"

var config map[string]string

// init() is called automatically — you can have multiple init() functions in one file
func init() {
	fmt.Println("[init 1] Loading config...")
	config = map[string]string{
		"host": "localhost",
		"port": "8080",
		"mode": "debug",
	}
}

func init() {
	fmt.Println("[init 2] Validating config...")
	if config["host"] == "" {
		panic("host is required in config")
	}
}

var computed int

func init() {
	fmt.Println("[init 3] Pre-computing values...")
	computed = len(config) * 10
}

func main() {
	fmt.Println("\n[main] Program starting")
	fmt.Println("Config:", config)
	fmt.Println("Computed:", computed)
}
