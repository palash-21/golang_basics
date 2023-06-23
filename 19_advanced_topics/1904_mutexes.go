/* In the previous example we saw how to manage simple counter state using atomic operations.
For more complex state we can use a mutex to safely access data across multiple goroutines.
*/

package main

import (
	"fmt"
	"sync"
)

/*
	Container holds a map of counters;

since we want to update it concurrently from multiple goroutines, we add a Mutex to synchronize access.
*/
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// Note that mutexes must not be copied, so if this struct is passed around, it should be done by pointer.
func (c *Container) inc(name string) {

	// Lock the mutex before accessing counters; unlock it at the end of the function using a defer statement.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	// Note that the zero value of a mutex is usable as-is, so no initialization is required here.
	c := Container{

		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// This function increments a named counter in a loop.
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	// Run several goroutines concurrently; note that they all access the same Container, and two of them access the same counter.
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	// Wait for the goroutines to finish
	wg.Wait()
	fmt.Println(c.counters)
}
