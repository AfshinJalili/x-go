// Package patterns demonstrates common channel patterns.
// This file implements the done channel pattern for cancellation.
package patterns

import (
    "fmt"
    "time"
)

// RunDoneChannel demonstrates the done channel pattern for graceful cancellation.
// Closing the done channel signals workers to stop processing.
func RunDoneChannel() {
    done := make(chan struct{})
    results := make(chan int)
    go worker1(done, results)

    go func() {
        for i := 0; i < 5; i++ {
            fmt.Println(<-results)
        }
    }()

    time.Sleep(200 * time.Millisecond)
    close(done)
}

func worker1(done <-chan struct{}, results chan<- int) {
    for i := 1; ; i++ {
        select {
        case <-done:
            fmt.Println("worker stopped")
            return
        case results <- i:
            time.Sleep(100 * time.Millisecond)
        }
    }
}