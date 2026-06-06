<div align="center">

# 🐹 go-learning

**Learn Go by reading and running 50 real examples one concept at a time**

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![Author](https://img.shields.io/badge/by-notcandy001-blueviolet)](https://github.com/notcandy001)

</div>


## Before you start, read this

This is not a book. There are no long explanations here.

Every lesson is a **folder**. Every folder has a **`.go` file**. That file *is* the lesson — the comments inside it explain everything line by line. You read the file, you run it, you mess with it, you move on.

That's the whole method. Simple.

## What's a `.go` file?

Every file in this repo ends with `.go` — that's the Go source file extension, the same way Python uses `.py` or JavaScript uses `.js`. When you see `main.go` inside a folder, **that file is the lesson**. Open it, read the comments at the top of each section, and follow along.

Here's what a typical lesson file looks like from the inside:

```go
// Variables hold data — Go is statically typed, so each variable has a fixed type
package main

import "fmt"

func main() {
    // Long form: declare type explicitly
    var name string = "Candy"

    // Short form: Go figures out the type for you
    age := 19

    fmt.Println(name, age)
}
```

The comments are written for you — they explain *what* is happening and *why*. Read them top to bottom before running anything.

---

## How to go through this repo

**Go in order.** The folders are numbered `00` through `49` for a reason. Each lesson builds on the ones before it. If you jump to `17_recursion` without reading `14_functions_basics` first, you'll be confused about things that were already explained.

You can skip things you already know, but when in doubt — start from `00`.

**For each folder, do this:**

1. `cd` into the folder
2. Open `main.go` and **read it top to bottom** — don't run it yet, just read
3. Look at the comments. Understand what each section is doing
4. Run it: `go run main.go`
5. Match the output to what you just read
6. **Change something** — swap a value, delete a line, add a print statement
7. Run it again. See what broke or changed
8. Move to the next folder

You don't need to memorize anything. The goal is to *feel* how Go works. The patterns repeat and you'll start recognizing them naturally.

---

## What's inside each folder?

Most folders look like this:

```
07_if_else/
├── main.go     ← the lesson — read this, run this, break this
└── go.mod      ← tells Go this folder is a self-contained program (ignore this file)
```

Two folders have an extra file — the testing lessons:

```
40_testing/
├── main.go        ← the functions you'll be testing
├── main_test.go   ← the test file itself (run with: go test -v)
└── go.mod

41_benchmarking/
├── main.go
├── bench_test.go  ← the benchmark file (run with: go test -bench=.)
└── go.mod
```

Everything else is just `main.go`. One file, one concept, one lesson.

---

## All 50 lessons at a glance

Read this table before you start so you know what's coming. Come back to it when you need to find something.

### 🟢 Level 1 — The Basics (00–09)
> Syntax, types, printing, input, math, strings, and control flow. Start here.

| Folder | What's inside the `.go` file |
|--------|------------------------------|
| `00_hello_world` | The smallest valid Go program. `package main`, `import`, `func main()`, `fmt.Println` |
| `01_variables_and_types` | How to store data. `var`, `:=`, `int`, `string`, `bool`, `float64` and zero values |
| `02_constants` | Values that never change. `const`, block declarations, package-level constants |
| `03_user_input` | Reading what the user types. `fmt.Scan` for one word, `bufio.Scanner` for full lines |
| `04_fmt_formatting` | Printing things properly. `%s`, `%d`, `%f`, `%v`, `%T`, `Printf`, `Sprintf` |
| `05_arithmetic` | Math in Go. `+`, `-`, `*`, `/`, `%` (remainder), `math.Sqrt`, `math.Pow` |
| `06_string_operations` | Working with text. `strings.Split`, `Join`, `Replace`, `Contains`, `TrimSpace` |
| `07_if_else` | Making decisions. `if`, `else if`, `else`, `&&` (and), `\|\|` (or), `!` (not) |
| `08_switch` | Cleaner branching. `switch`, multiple values per case, expressionless switch, type switch |
| `09_for_loops` | Go's only loop keyword. `for`, `range`, iterating over slices, maps, and strings |

### 🔵 Level 2 — Data Structures & Functions (10–19)
> Go's core data types and how to organize code into reusable pieces.

| Folder | What's inside the `.go` file |
|--------|------------------------------|
| `10_while_style_loops` | `for` used as a `while` loop. `break`, `continue`, and infinite loops |
| `11_arrays` | Fixed-size lists. Why arrays exist, value semantics, why you'll prefer slices |
| `12_slices` | Dynamic lists — the most important data structure in Go. `append`, `make`, `copy`, 2D slices |
| `13_maps` | Key-value storage. `map[string]int`, adding, deleting, checking if a key exists |
| `14_functions_basics` | Writing functions. Parameters, return values, functions stored in variables |
| `15_multiple_return_values` | Go functions can return two things at once. The `(value, error)` pattern you'll see everywhere |
| `16_variadic_functions` | Functions that accept any number of arguments. `...int`, spreading a slice |
| `17_recursion` | Functions that call themselves. Base case, factorial, Fibonacci, recursive sum |
| `18_closures` | Functions that remember variables from outside themselves. Counters, adders |
| `19_pointers` | Memory addresses. `&` (address of), `*` (value at that address), passing by reference |

### 🟡 Level 3 — Types, Interfaces & Concurrency (20–29)
> Go's type system, how to model real things, and how to run code concurrently.

| Folder | What's inside the `.go` file |
|--------|------------------------------|
| `20_structs` | Grouping related data into a custom type. Struct literals, nested structs, pointer structs |
| `21_methods` | Functions attached to a struct. Value vs pointer receivers, the `String()` method |
| `22_interfaces` | Define behaviour without caring about the concrete type. Polymorphism in Go |
| `23_embedding` | One struct using another's fields and methods. Go's alternative to class inheritance |
| `24_error_handling` | Errors are values, not exceptions. `errors.New`, `fmt.Errorf`, `%w` for wrapping |
| `25_custom_errors` | Building error types with extra fields. `errors.As` to inspect them |
| `26_goroutines` | Running code concurrently with the `go` keyword. Lightweight, not OS threads |
| `27_channels` | Goroutines communicating through pipes. `chan`, send `<-`, receive `<-`, buffered channels |
| `28_select_statement` | Waiting on multiple channels at once. Timeout pattern, non-blocking with `default` |
| `29_waitgroup` | Waiting for a group of goroutines to finish. `sync.WaitGroup` — `Add`, `Done`, `Wait` |

### 🔴 Level 4 — The Real World (30–49)
> Files, JSON, HTTP, CLI tools, testing, algorithms, and a full working project.

| Folder | What's inside the `.go` file |
|--------|------------------------------|
| `30_mutex` | Preventing race conditions. `sync.Mutex` to lock shared data, `sync.RWMutex` for read-heavy access |
| `31_file_read_write` | Reading and writing files on disk. `os.ReadFile`, `os.WriteFile`, line-by-line with `bufio` |
| `32_json_encode_decode` | Go structs ↔ JSON. `json.Marshal`, `json.Unmarshal`, struct tags like `` `json:"name"` `` |
| `33_http_get` | Making HTTP requests. GET a URL, read the body, decode a JSON response |
| `34_http_server` | Building an HTTP server. Handlers, routes, reading query params |
| `35_os_args` | Command-line arguments. Raw `os.Args` and structured flags with the `flag` package |
| `36_env_variables` | Reading config from the shell environment. `os.Getenv`, `os.Setenv`, `os.LookupEnv` |
| `37_defer_panic_recover` | Code that always runs on exit — `defer`. Crashing on purpose — `panic`. Catching crashes — `recover` |
| `38_type_assertions` | Getting the real type out of an interface. `i.(string)`, safe two-value form, type switch |
| `39_generics` | Functions that work with any type. `[T comparable]`, type constraints, a generic Stack |
| `40_testing` | Writing and running tests. `go test`, `testing.T`, table-driven tests (the Go-idiomatic way) |
| `41_benchmarking` | Measuring how fast your code is. `testing.B`, `go test -bench=.`, `-benchmem` |
| `42_build_tags` | Compiling code only for certain platforms. `//go:build linux`, `runtime.GOOS` |
| `43_init_function` | Code that runs automatically before `main()`. `init()`, multiple inits, execution order |
| `44_blank_identifier` | Discarding values you don't need. `_` — when Go forces you to handle something you don't want |
| `45_iota_enums` | Auto-incrementing constants. `iota`, enum patterns, byte sizes (`KB`, `MB`), bit flags |
| `46_linked_list` | Building a classic data structure from scratch. Nodes, pointers, append, delete, print |
| `47_binary_search` | Finding things fast in a sorted list. O(log n), iterative version and recursive version |
| `48_sorting` | Sorting slices and custom types. `sort.Ints`, `sort.Slice` with a custom comparator |
| `49_mini_project_todo` | **A real CLI todo app.** Structs, methods, JSON file persistence, bufio input, switch dispatch — everything combined |

---

## Installing Go

1. Download the installer for your OS from [https://go.dev/dl](https://go.dev/dl)
2. Install it, then open a terminal and check it worked:

```bash
go version
# go version go1.21.x linux/amd64
```

### Editor setup (pick one)

- **VS Code** — install the Go extension: `ext install golang.go`
- **Neovim** — use gopls as your LSP
- **GoLand** — JetBrains IDE, paid but excellent

---

## Running any lesson

```bash
cd 00_hello_world
go run main.go
```

That's it. No install step, no build script, no configuration. Just `go run main.go` from inside any folder.

For the two testing lessons:

```bash
cd 40_testing
go test -v          # runs the tests, shows pass/fail per test

cd 41_benchmarking
go test -bench=. -benchmem    # runs benchmarks, shows ns/op and memory
```

---

## Useful commands to know

```bash
go run main.go       # run a .go file directly without compiling first
go build             # compile the current folder into a binary
go fmt main.go       # auto-format your code (run this often)
go vet main.go       # catch common mistakes before they bite you
go doc fmt.Println   # read the docs for any function right in your terminal
go test -v           # run tests with verbose output
go test -bench=.     # run benchmarks
```

---

## Tips

**Read the file before running it.** The comments are the lesson. If you just run it without reading, you're skipping the whole point.

**Break things on purpose.** Delete a line. Change a type. Remove a `return`. See what error Go gives you. The error messages are surprisingly helpful and learning to read them fast is a skill.

**Go is strict.** If you declare a variable and don't use it, it won't compile. If you import a package and don't use it, it won't compile. This feels annoying at first but it keeps code clean.

**Errors are not exceptions.** You will see `if err != nil { ... }` constantly. That's intentional. Go makes you handle errors where they happen instead of letting them bubble up silently.

**The standard library is huge and good.** Before looking for a third-party package, check if `strings`, `sort`, `os`, `net/http`, or `encoding/json` already does what you need. It probably does.


## 📝 License

MIT License use it however you want.

<div align="center">
Made for learning by <a href="https://github.com/notcandy001">notcandy001</a>
</div>
