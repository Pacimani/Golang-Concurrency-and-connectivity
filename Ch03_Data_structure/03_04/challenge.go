package main

import (
	"fmt"
	"net/http"
)

// write a function that returns content-type header

func main() {

	ctype, err := getContentType("https://linkedin.com")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(ctype)
}

func getContentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close() // make sure to close the body

	ctype := resp.Header.Get("Content-Type")

	if ctype == "" {
		return "", fmt.Errorf("Content-Type not found")
	}
	return ctype, nil
}
