// Environment variables are key-value pairs set in the OS shell — useful for config
package main

import (
	"fmt"
	"os"
)

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func main() {
	// os.Getenv returns the value or empty string if not set
	home := os.Getenv("HOME")
	fmt.Println("HOME:", home)

	// os.LookupEnv returns (value, found) — lets you distinguish empty vs missing
	path, ok := os.LookupEnv("PATH")
	if ok {
		fmt.Println("PATH is set (first 50 chars):", path[:min(50, len(path))])
	}

	// Set an env var for this process
	os.Setenv("APP_MODE", "development")
	fmt.Println("APP_MODE:", os.Getenv("APP_MODE"))

	// Use a fallback for optional config
	port := getEnv("PORT", "8080")
	dbHost := getEnv("DB_HOST", "localhost")
	fmt.Printf("Server will run on :%s\n", port)
	fmt.Printf("DB host: %s\n", dbHost)

	// Unset a variable
	os.Unsetenv("APP_MODE")
	fmt.Println("After Unsetenv, APP_MODE:", os.Getenv("APP_MODE"))

	// List all environment variables
	// envs := os.Environ()
	// for _, e := range envs { fmt.Println(e) }
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
