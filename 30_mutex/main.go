// sync.Mutex prevents race conditions when multiple goroutines access shared data
package main

import (
	"fmt"
	"sync"
)

// SafeCounter is safe to use from multiple goroutines concurrently
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()   // only one goroutine can be here at a time
	c.value++
	c.mu.Unlock() // always unlock — or use defer
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock() // defer ensures unlock even if we panic
	return c.value
}

// RWMutex allows multiple concurrent readers but only one writer
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *Cache) Set(key, val string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = val
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock() // multiple goroutines can RLock at the same time
	defer c.mu.RUnlock()
	v, ok := c.data[key]
	return v, ok
}

func main() {
	counter := &SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()
	fmt.Println("Final counter:", counter.Value()) // should always be 1000

	cache := &Cache{data: make(map[string]string)}
	cache.Set("lang", "Go")
	cache.Set("version", "1.22")

	if v, ok := cache.Get("lang"); ok {
		fmt.Println("Cached lang:", v)
	}
}
