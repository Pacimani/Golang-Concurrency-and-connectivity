package main

// go concurrency is called goroutines
// they can spin several go processes at the same time on the same machine

import (
	"fmt"
	"net/http"
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

func main() {

	urls := []string{"https://httpbin.com",
		"https://api.github.com", "https://golang.org"}

	start := time.Now()

	siteSerial(urls)

	fmt.Printf("Duration: %v\n", time.Since(start))
}
