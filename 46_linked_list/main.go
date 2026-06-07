// Implementing a singly linked list from scratch — a classic data structure
package main

import "fmt"

// Node holds a value and a pointer to the next node
type Node struct {
	Value int
	Next  *Node
}

// LinkedList keeps track of the head node and its length
type LinkedList struct {
	Head   *Node
	Length int
}

// Append adds a node to the end of the list
func (l *LinkedList) Append(val int) {
	newNode := &Node{Value: val}
	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	l.Length++
}

// Prepend adds a node to the front
func (l *LinkedList) Prepend(val int) {
	newNode := &Node{Value: val, Next: l.Head}
	l.Head = newNode
	l.Length++
}

// Delete removes the first node with the given value
func (l *LinkedList) Delete(val int) {
	if l.Head == nil {
		return
	}
	if l.Head.Value == val {
		l.Head = l.Head.Next
		l.Length--
		return
	}
	current := l.Head
	for current.Next != nil {
		if current.Next.Value == val {
			current.Next = current.Next.Next
			l.Length--
			return
		}
		current = current.Next
	}
}

// Print displays the list
func (l *LinkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Printf("%d", current.Value)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Printf(" (len=%d)\n", l.Length)
}

func main() {
	list := &LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Prepend(0)
	list.Print() // 0 -> 1 -> 2 -> 3

	list.Delete(2)
	list.Print() // 0 -> 1 -> 3
}
