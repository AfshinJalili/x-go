// Package patterns demonstrates common channel patterns.
// This file implements the generator pattern using channels.
package patterns

import "fmt"

// RunGenerator demonstrates the generator pattern.
// A generator function returns a channel that produces a sequence of values.
func RunGenerator() {
    fib := fibonacci(10)
    for num := range fib {
        fmt.Println(num)
    }
}

func fibonacci(n int) <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        a, b := 0, 1
        for i := 0; i < n; i++ {
            ch <- a
            a, b = b, a+b
        }
    }()
    return ch
}