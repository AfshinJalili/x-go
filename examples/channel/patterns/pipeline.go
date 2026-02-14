// Package patterns demonstrates common channel patterns.
// This file implements the pipeline pattern for processing data through stages.
package patterns

import "fmt"

// RunPipeline demonstrates the pipeline pattern.
// Data flows through multiple stages: generate -> square -> sum.
func RunPipeline() {
    numbers := generate(1, 2, 3, 4, 5, 6)
    squared := square(numbers)
    result := sum(squared)
    fmt.Println("sum power 2:", <-result)
}

func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func sum(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        total := 0
        for n := range in {
            total += n
        }
        out <- total
        close(out)
    }()
    return out
}