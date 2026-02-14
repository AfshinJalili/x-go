// Package selectpkg demonstrates the select statement for channel operations.
// This file shows basic select usage with multiple channels.
package selectpkg

import (
    "fmt"
    "time"
)

// UsingSelect demonstrates basic select statement usage.
// Select waits on multiple channels and processes the first one that becomes ready.
func UsingSelect() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(100 * time.Millisecond)
        ch1 <- "message from ch1"
    }()
    
    go func() {
        time.Sleep(50 * time.Millisecond)
        ch2 <- "message from ch2"
    }()
    
    for range 2 {
        select {
        case msg1 := <-ch1:
            fmt.Println("from ch1:", msg1)
        case msg2 := <-ch2:
            fmt.Println("from ch2:", msg2)
        }
    }
}
