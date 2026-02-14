// Package selectpkg demonstrates the select statement for channel operations.
// This file shows using select with time.After for timeout handling.
package selectpkg

import (
    "fmt"
    "time"
)

// SelectTimeout demonstrates using select with time.After to implement timeouts.
// If the channel doesn't receive a value within the timeout period, the timeout case executes.
func SelectTimeout() {
    ch := make(chan string)
    go func() {
        time.Sleep(2 * time.Second)
        ch <- "answer"
    }()
    
    select {
    case msg := <-ch:
        fmt.Println("received:", msg)
    case <-time.After(3 * time.Second):
        fmt.Println("Timeout!")
    }
}
