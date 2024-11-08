/*
	Calculate total download size for NYC taxi data for 2020

For each month, we have two files: green and yellow. For example:

	https://s3.amazonaws.com/nyc-tlc/trip+data/green_tripdata_2020-03.csv
	https://s3.amazonaws.com/nyc-tlc/trip+data/yellow_tripdata_2020-03.csv

Turn the below sequential algorithm to a concurrent one using goroutine per file.
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// url template
var (
	urlTemplate = "https://s3.amazonaws.com/nyc-tlc/trip+data/%s_tripdata_2020-%02d.csv"
	colors      = []string{"green", "yellow"}
)

// return the size of the file in the header without returning the whole file
func downloadSize(url string) (int, error) {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(resp.Status)
	}

	return strconv.Atoi(resp.Header.Get("Content-Length"))
}

type result struct {
	url  string
	size int
	err  error
}

func sizeWorker(url string, ch chan result) {

	fmt.Printf("Downloading %s\n", url)

	res := result{url: url}
	res.size, res.err = downloadSize(url)
	ch <- res

}

func main() {
	start := time.Now()

	size := 0

	ch := make(chan result)
	for _, color := range colors {
		for month := 1; month <= 12; month++ {
			url := fmt.Sprintf(urlTemplate, color, month)

			go sizeWorker(url, ch)
		}
	}

	for i := 0; i < len(colors)*12; i++ {
		res := <-ch

		if res.err != nil {
			log.Fatal(res.err)
		}
		size += res.size
	}

	fmt.Printf("Duration: %v and total size is %d\n", time.Since(start), size)

}
