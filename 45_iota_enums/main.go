// iota is a special constant that auto-increments in a const block — perfect for enums
package main

import "fmt"

// Weekday enum — iota starts at 0 and increments by 1 for each const
type Weekday int

const (
	Sunday Weekday = iota // 0
	Monday                // 1
	Tuesday               // 2
	Wednesday             // 3
	Thursday              // 4
	Friday                // 5
	Saturday              // 6
)

func (d Weekday) String() string {
	names := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	if int(d) < len(names) {
		return names[d]
	}
	return "Unknown"
}

// Byte size enum using iota with an expression
type ByteSize float64

const (
	_           = iota             // discard first value (0)
	KB ByteSize = 1 << (10 * iota) // 1 << 10 = 1024
	MB                             // 1 << 20
	GB                             // 1 << 30
	TB                             // 1 << 40
)

// Bit flags — combine with |
type Permission uint

const (
	Read    Permission = 1 << iota // 1
	Write                          // 2
	Execute                        // 4
)

func main() {
	fmt.Println("Days:", Sunday, Monday, Friday, Saturday)

	fmt.Printf("KB = %.0f bytes\n", float64(KB))
	fmt.Printf("MB = %.0f bytes\n", float64(MB))
	fmt.Printf("GB = %.0f bytes\n", float64(GB))

	perms := Read | Write
	fmt.Printf("Perms: %b (Read=%v, Write=%v, Execute=%v)\n",
		perms,
		perms&Read != 0,
		perms&Write != 0,
		perms&Execute != 0,
	)
}
