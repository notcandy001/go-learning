// Channels are typed pipes that goroutines use to communicate and sync with each other
package main

import "fmt"

// Worker sends results back through a channel
func square(n int, ch chan int) {
	ch <- n * n // send value into channel
}

func producer(ch chan string) {
	items := []string{"apple", "banana", "cherry"}
	for _, item := range items {
		ch <- item // send each item
	}
	close(ch) // close signals "no more values"
}

func main() {
	// Unbuffered channel — send blocks until someone receives
	ch := make(chan int)
	go square(7, ch)
	result := <-ch // receive from channel
	fmt.Println("7^2 =", result)

	// Buffered channel — holds up to N values without blocking
	buf := make(chan string, 3)
	buf <- "one"
	buf <- "two"
	buf <- "three"
	fmt.Println(<-buf) // one
	fmt.Println(<-buf) // two
	fmt.Println(<-buf) // three

	// Range over a channel — reads until the channel is closed
	strCh := make(chan string, 5)
	go producer(strCh)
	for item := range strCh {
		fmt.Println("Got:", item)
	}

	// Fan-out: launch multiple goroutines, collect results
	results := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		go square(i, results)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-results)
	}
}
