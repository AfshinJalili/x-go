// Package patterns demonstrates common channel patterns.
// This file implements a semaphore using a buffered channel to limit concurrency.
package patterns

import (
    "fmt"
    "time"
)

// RunSemaphore demonstrates using a buffered channel as a semaphore.
// The semaphore limits the number of concurrent goroutines to the channel's capacity.
func RunSemaphore() {
    semaphore := make(chan struct{}, 3)
    for i := 1; i <= 10; i++ {
        semaphore <- struct{}{}
        go func(id int) {
            defer func() { <-semaphore }()
            fmt.Printf("job %d started\n", id)
            time.Sleep(time.Second)
            fmt.Printf("job %d done\n", id)
        }(i)
    }
    time.Sleep(5 * time.Second)
}