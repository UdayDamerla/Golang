// Go routines are lightweight threads managed by the Go runtime.
// They are a key feature of Go that allows concurrent programming.
// In this tutorial, we explore how to create and manage goroutines in Go.
//
// Notes:
// - Goroutines are not the same as OS threads. They are much lighter and efficient.
// - We use lock/unlock so only one goroutine can access critical sections at a time.
// - defer ensures unlock is called even if function returns early.
// - Without proper synchronization, race conditions can cause unpredictable behavior.

package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	// RWMutex allows multiple readers or a single writer.
	mu    sync.RWMutex
	value int
}

func (c *SafeCounter) Increment(writerID int) {
	// Exclusive lock for write operation.
	c.mu.Lock()
	c.value++
	current := c.value
	c.mu.Unlock()

	fmt.Printf("Writer %d incremented counter to %d\n", writerID, current)
}

func (c *SafeCounter) Read(readerID int) int {
	// Shared lock for read operation.
	c.mu.RLock()
	current := c.value
	c.mu.RUnlock()

	fmt.Printf("Reader %d read counter as %d\n", readerID, current)
	return current
}

func sayHello(id int, wg *sync.WaitGroup, printMu *sync.Mutex) {
	// Signal completion when this goroutine exits.
	defer wg.Done()

	// Mutex protects this print block as a critical section.
	printMu.Lock()
	fmt.Printf("Hello from goroutine %d\n", id)
	printMu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var printMu sync.Mutex

	fmt.Println("=== Part 1: Mutex for critical section ===")
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go sayHello(i, &wg, &printMu)
	}
	wg.Wait()

	fmt.Println("\n=== Part 2: RWMutex for shared data ===")
	counter := &SafeCounter{}

	writers := 3
	readers := 4
	incrementsPerWriter := 2
	readsPerReader := 3

	for w := 1; w <= writers; w++ {
		wg.Add(1)
		go func(writerID int) {
			defer wg.Done()
			for i := 0; i < incrementsPerWriter; i++ {
				counter.Increment(writerID)
				time.Sleep(50 * time.Millisecond)
			}
		}(w)
	}

	for r := 1; r <= readers; r++ {
		wg.Add(1)
		go func(readerID int) {
			defer wg.Done()
			for i := 0; i < readsPerReader; i++ {
				counter.Read(readerID)
				time.Sleep(30 * time.Millisecond)
			}
		}(r)
	}

	wg.Wait()
	fmt.Printf("\nFinal counter value: %d\n", counter.Read(0))
	fmt.Println("Main function completed.")
}
