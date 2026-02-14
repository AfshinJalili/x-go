package examples

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}


func MutexLock() {
	counter := &Counter{}
	var wg sync.WaitGroup

	// 1000 concurrent Goroutine
	for range 1000 {
		wg.Go(func() {
			counter.Increment()
		})
	}

	wg.Wait()
	fmt.Println("final value: ", counter.Value())
}