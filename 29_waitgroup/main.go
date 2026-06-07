// sync.WaitGroup lets you wait for a collection of goroutines to finish
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrement counter when this function returns — always paired with Add
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // increment counter before starting goroutine
		go worker(i, &wg)
	}

	wg.Wait() // block until all workers call wg.Done()
	fmt.Println("All workers finished")

	// Collecting results safely using a mutex-protected slice
	var mu sync.Mutex
	var results []int

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mu.Lock()
			results = append(results, n*n)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println("Squares:", results)
}
