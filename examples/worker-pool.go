package examples 

import (
	"fmt"
	"time"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d ready to process jobs\n", id) 

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job) 
		time.Sleep(500 * time.Millisecond) 
		results <- job * 2
	}
}


func WorkerPool() {
	numJobs := 10
	numWorkers := 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	fmt.Println("waiting for 5 sec")
	time.Sleep(time.Second * 5)
	

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

