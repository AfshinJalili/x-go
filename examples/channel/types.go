// Package channel demonstrates basic channel types and operations.
// This file shows different channel types: unbuffered, buffered, send-only, and receive-only channels.
package channel

import (
    "fmt"
    "time"
)

// DifferentTypesOfChannel demonstrates unbuffered channels, buffered channels,
// and directional channel types (send-only and receive-only).
func DifferentTypesOfChannel() {
    ch := make(chan int)
    
    bufferedCh := make(chan string, 3)
    
    var sendOnly chan<- int = ch
    var receiveOnly <-chan int = ch
    
    fmt.Printf("ch type: %T\n", ch)
    fmt.Printf("bufferedCh type: %T\n", bufferedCh)
    fmt.Println(sendOnly, receiveOnly)
    
    go func() {
        ch <- 1
    }()
    fmt.Println(<-ch)
    
    go func() {
        time.Sleep(time.Second * 5)
        ch <- 42
    }()
    fmt.Println("received:", <-ch)
    
    bufferedCh <- "not blocked"
    bufferedCh <- "not blocked yet"
    bufferedCh <- "still... not blocked"
    fmt.Println(<-bufferedCh)
    fmt.Println(<-bufferedCh)
    fmt.Println(<-bufferedCh)
}
