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
	// create a new http request
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return 0, err
	}

	// does the head request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	// the function should return the size of the file

	// check the status
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("bad status: %s", resp.Status)
	}

	return strconv.Atoi(resp.Header.Get("Content-Length"))
}

func main() {
	start := time.Now()

	size := 0
	for _, color := range colors {
		for month := 1; month <= 12; month++ {
			url := fmt.Sprintf(urlTemplate, color, month)
			n, err := downloadSize(url)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s %s: %d\n", color, time.Now().Format("2006-01-02 15:04:05"), n)

			size += n
		}
	}
	fmt.Printf("Duration: %v and total size is %d\n", time.Since(start), size)
}
