// Build tags (constraints) control which files are compiled based on OS, arch, or custom tags
package main

import (
	"fmt"
	"runtime"
)

// You can also use //go:build linux (or //go:build windows) at the top of a file
// to include/exclude it from compilation on specific platforms.

func main() {
	// runtime package lets you check the OS at runtime (simpler than build tags for simple checks)
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("CPUs:", runtime.NumCPU())

	switch runtime.GOOS {
	case "linux":
		fmt.Println("Running on Linux — great for ricing!")
	case "darwin":
		fmt.Println("Running on macOS")
	case "windows":
		fmt.Println("Running on Windows")
	default:
		fmt.Println("Unknown OS")
	}

	// To use build tags, put at the top of a file (before package declaration):
	// //go:build linux
	// // +build linux   (old syntax, keep for compatibility)
	//
	// Then build with: go build -tags yourtag
}
