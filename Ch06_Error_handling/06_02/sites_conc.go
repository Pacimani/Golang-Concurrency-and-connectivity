package main

// go concurrency is called goroutines
// they can spin several go processes at the same time on the same machine

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {

		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close() // make sure to close the body

	ctype := resp.Header.Get("Content-Type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func siteSerial(url []string) {
	for _, url := range url {
		returnType(url)
	}
}

// a concurent version of siteSerial// A concurrent version of siteSerial
func sitesConcurrent(urls []string) {
	// Create a WaitGroup to track the number of goroutines
	var wg sync.WaitGroup

	// Iterate over each URL
	for _, url := range urls {
		// Increment the WaitGroup counter for each goroutine
		wg.Add(1)
		// Start a new goroutine for each URL
		go func(url string) {
			// Fetch the content type of the URL
			returnType(url)
			// Decrement the WaitGroup counter when done
			wg.Done()
		}(url)
	}
	// Wait for all goroutines to finish
	wg.Wait()
}

func main() {

	urls := []string{"https://httpbin.com",
		"https://api.github.com", "https://golang.org"}

	start := time.Now()

	siteSerial(urls)

	fmt.Printf("Duration: %v\n", time.Since(start))

	start = time.Now()

	sitesConcurrent(urls)

	fmt.Printf("Duration concurrent: %v\n", time.Since(start))
}
