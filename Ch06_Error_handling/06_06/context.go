package main

import (
	"context"
	"log"
	"time"
)

type Bid struct {
	AdURL string
	Price float64
}

func bestBid(url string) Bid {
	time.Sleep(2 * time.Second)
	return Bid{
		AdURL: "https://github.com",
		Price: 100.00,
	}
}

var defaultBid = Bid{
	AdURL: "https://github.com",
	Price: 1009.00,
}

func findBid(ctx context.Context, url string) Bid {

	ch := make(chan Bid, 1) // buffered channel to avoid goroutine leaks

	go func() {

		ch <- bestBid(url)
	}()
	select {
	case bid := <-ch:
		return bid
		// if the algorithm is not finished in the allocated time, we return the default bid
	case <-ctx.Done():
		log.Println("findBid: context canceled")
		return defaultBid
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	url := "https://http.cat/418"

	bid := findBid(ctx, url)
	log.Println(bid)
	log.Println("end")
}
