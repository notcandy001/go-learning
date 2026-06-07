// select waits on multiple channel operations — like a switch for channels
package main

import (
	"fmt"
	"time"
)

func fastTask(ch chan string) {
	time.Sleep(100 * time.Millisecond)
	ch <- "fast task done"
}

func slowTask(ch chan string) {
	time.Sleep(500 * time.Millisecond)
	ch <- "slow task done"
}

func main() {
	fast := make(chan string, 1)
	slow := make(chan string, 1)

	go fastTask(fast)
	go slowTask(slow)

	// select picks whichever channel is ready first
	for i := 0; i < 2; i++ {
		select {
		case msg := <-fast:
			fmt.Println("Fast:", msg)
		case msg := <-slow:
			fmt.Println("Slow:", msg)
		}
	}

	// select with default — non-blocking check
	ch := make(chan int, 1)
	select {
	case v := <-ch:
		fmt.Println("Got:", v)
	default:
		fmt.Println("No value ready — continuing")
	}

	// Timeout pattern using time.After
	result := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		result <- "too late"
	}()

	select {
	case r := <-result:
		fmt.Println("Got:", r)
	case <-time.After(300 * time.Millisecond):
		fmt.Println("Timed out!")
	}
}
