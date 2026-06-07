// os.Args gives you command-line arguments — the flag package makes them easier to parse
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// os.Args[0] is always the program name, [1:] are the actual args
	fmt.Println("Program:", os.Args[0])
	fmt.Println("All args:", os.Args)

	if len(os.Args) > 1 {
		fmt.Println("First arg:", os.Args[1])
	}

	// flag package for structured CLI flags
	// Run as: go run main.go -name Candy -count 3
	name := flag.String("name", "World", "who to greet")
	count := flag.Int("count", 1, "how many times to greet")
	verbose := flag.Bool("verbose", false, "enable verbose output")

	flag.Parse() // must call this before using flag values

	if *verbose {
		fmt.Println("Verbose mode enabled")
		fmt.Printf("Greeting %q %d time(s)\n", *name, *count)
	}

	for i := 0; i < *count; i++ {
		fmt.Printf("Hello, %s!\n", *name)
	}

	// flag.Args() returns non-flag arguments after --
	remaining := flag.Args()
	if len(remaining) > 0 {
		fmt.Println("Extra args:", remaining)
	}
}
