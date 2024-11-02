package main

import (
	"fmt"
	"net/http"
)

func returnType(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("%s -> Error: %s\n", url, err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("Content-Type")
	out <- fmt.Sprintf("%s -> %s\n", url, ctype)
}

func main() {

	urls := []string{"https://httpbin.com",
		"https://api.github.com", "https://golang.org"}

	ch := make(chan string)
	for _, url := range urls {
		go returnType(url, ch)
	}

	for range urls {
		fmt.Printf("Channel: %s\n", <-ch)
	}
}
