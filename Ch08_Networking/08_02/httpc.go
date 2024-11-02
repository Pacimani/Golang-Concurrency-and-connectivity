package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Job is a job description
type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func main() {

	// get request

	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("Error: cant call https://httpbin.org/get -> %s", err)
	}
	defer resp.Body.Close()

	// copy to std out whatever is in the response body
	io.Copy(os.Stdout, resp.Body)

	fmt.Println("----------")

	// Post a job
	job := &Job{
		User:   "Pacifique Rukiza",
		Action: "punch",
		Count:  12,
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(job); err != nil {
		log.Fatalf("Error: cant encode job -> %s", err)
	}

	resp, err = http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatalf("Error: cant post to https://httpbin.org/post -> %s", err)
	}
	defer resp.Body.Close()

	// copy to std out whatever is in the response body
	io.Copy(os.Stdout, resp.Body)

}
