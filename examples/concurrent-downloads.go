package examples 

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL string
	Size int
	Error error
}

func fetchURL(url string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		results <- Result{URL: url, Error: err}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		results <- Result{URL: url, Error: err}
		return
	}

	elapsed := time.Since(start)
	fmt.Printf("%s: %d bytes in %v\n", url, len(body), elapsed)

	results <- Result{URL: url, Size: len(body)}
}


func ConcurrentDownloads() {
	urls := []string{
		"https://downloadly.ir",
        "https://google.com",
        "https://github.com",
	}

	results := make(chan Result, len(urls))
	var wg sync.WaitGroup

	start := time.Now()

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	totalSize := 0
	for result := range results {
		if result.Error == nil {
			totalSize += result.Size
		}
	}

	fmt.Printf("\nTotal: %d bytes in %v\n", totalSize, time.Since(start))
}