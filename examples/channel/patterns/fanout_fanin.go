// Package patterns demonstrates common channel patterns.
// This file implements the fan-out/fan-in pattern for parallel processing.
package patterns

import (
    "fmt"
    "sync"
)

// RunFanoutFanin demonstrates the fan-out/fan-in pattern.
// Multiple workers process jobs from a single input channel and send results to a single output channel.
func RunFanoutFanin() {
    numJobs := 10
    numWorkers := 3
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    var wg sync.WaitGroup
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            worker(id, jobs, results)
        }(w)
    }

    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    go func() {
        wg.Wait()
        close(results)
    }()

    for result := range results {
        fmt.Println("result:", result)
    }
}

func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        results <- job * 2
    }
}