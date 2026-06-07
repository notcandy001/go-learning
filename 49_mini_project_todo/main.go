// Mini project: a fully functional CLI todo app — combining everything learned so far
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Todo represents a single task
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

// TodoList manages all todos and handles persistence
type TodoList struct {
	todos    []Todo
	filename string
	nextID   int
}

func NewTodoList(filename string) *TodoList {
	tl := &TodoList{filename: filename, nextID: 1}
	tl.load()
	return tl
}

func (tl *TodoList) Add(title string) {
	todo := Todo{ID: tl.nextID, Title: title, Done: false, CreatedAt: time.Now()}
	tl.todos = append(tl.todos, todo)
	tl.nextID++
	tl.save()
	fmt.Printf("Added [%d]: %s\n", todo.ID, todo.Title)
}

func (tl *TodoList) Complete(id int) {
	for i, t := range tl.todos {
		if t.ID == id {
			tl.todos[i].Done = true
			tl.save()
			fmt.Printf("Completed: %s\n", t.Title)
			return
		}
	}
	fmt.Println("Todo not found:", id)
}

func (tl *TodoList) Delete(id int) {
	for i, t := range tl.todos {
		if t.ID == id {
			tl.todos = append(tl.todos[:i], tl.todos[i+1:]...)
			tl.save()
			fmt.Printf("Deleted: %s\n", t.Title)
			return
		}
	}
	fmt.Println("Todo not found:", id)
}

func (tl *TodoList) List() {
	if len(tl.todos) == 0 {
		fmt.Println("No todos yet! Type 'add <task>'")
		return
	}
	fmt.Println("─────────────────────────────")
	for _, t := range tl.todos {
		status := "[ ]"
		if t.Done {
			status = "[✓]"
		}
		fmt.Printf("%s %d. %s\n", status, t.ID, t.Title)
	}
	fmt.Println("─────────────────────────────")
}

func (tl *TodoList) save() {
	data, _ := json.MarshalIndent(tl.todos, "", "  ")
	os.WriteFile(tl.filename, data, 0644)
}

func (tl *TodoList) load() {
	data, err := os.ReadFile(tl.filename)
	if err != nil {
		return // file doesn't exist yet — that's fine
	}
	json.Unmarshal(data, &tl.todos)
	// Recalculate nextID
	for _, t := range tl.todos {
		if t.ID >= tl.nextID {
			tl.nextID = t.ID + 1
		}
	}
}

func printHelp() {
	fmt.Println("Commands:")
	fmt.Println("  list              — show all todos")
	fmt.Println("  add <title>       — add a new todo")
	fmt.Println("  done <id>         — mark a todo as complete")
	fmt.Println("  delete <id>       — delete a todo")
	fmt.Println("  help              — show this help")
	fmt.Println("  quit              — exit")
}

func main() {
	tl := NewTodoList("/tmp/go_todos.json")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("┌─────────────────────────┐")
	fmt.Println("│   Go Todo App  v1.0     │")
	fmt.Println("│   Type 'help' to start  │")
	fmt.Println("└─────────────────────────┘")

	for {
		fmt.Print("\n> ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		parts := strings.SplitN(line, " ", 2)
		cmd := strings.ToLower(parts[0])
		arg := ""
		if len(parts) > 1 {
			arg = parts[1]
		}

		switch cmd {
		case "list", "ls":
			tl.List()
		case "add":
			if arg == "" {
				fmt.Println("Usage: add <title>")
			} else {
				tl.Add(arg)
			}
		case "done":
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Usage: done <id>")
			} else {
				tl.Complete(id)
			}
		case "delete", "del", "rm":
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Usage: delete <id>")
			} else {
				tl.Delete(id)
			}
		case "help", "h", "?":
			printHelp()
		case "quit", "exit", "q":
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Printf("Unknown command: %q — type 'help'\n", cmd)
		}
	}
}
