// Package channel demonstrates basic channel types and operations.
// This file shows how to properly close channels and detect channel closure.
package channel

import "fmt"

func producer(ch chan int) {
    for i := 1; i <= 5; i++ {
        ch <- i
    }
    close(ch)
}

// ClosingChannel demonstrates two methods for handling channel closure:
// 1. Using range to iterate over channel values until closed
// 2. Using the ok idiom to check if a channel is closed
func ClosingChannel() {
    ch := make(chan int)
    go producer(ch)
    
    for value := range ch {
        fmt.Println(value)
    }
    
    ch2 := make(chan int)
    go func() {
        defer close(ch2)
        ch2 <- 42
    }()
    
    value, ok := <-ch2
    fmt.Println(value, ok)
    value, ok = <-ch2
    fmt.Println(value, ok)
}
