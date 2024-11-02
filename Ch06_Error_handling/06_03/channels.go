package main

import (
	"fmt"
	"time"
)

// the prefer way to comunicate betweeen goroutines is via channels
// you can think of a channel as one unidirectional pipe with a specific type data type
// such as int, string, struct, etc
// the two important things about channels are the directionality and the type of data you are sending and receiving a value. if there is no goroutine on the other side, the channel will block.
// this works for both sender and receiver goroutine

func main() {

	start := time.Now()
	ch := make(chan int)

	go func() {
		ch <- 353 // send number to the channel
	}()

	x := <-ch // receive
	fmt.Printf("x = %d\n", x)
	fmt.Printf("Duration: %v\n", time.Since(start))
	println("-------")

	// send multiple values
	start = time.Now()

	count := 4
	go func() {
		for i := 0; i < count; i++ {
			fmt.Printf("Sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}

	}()

	// receive multiple values
	for i := 0; i < count; i++ {
		val := <-ch
		fmt.Printf("Received %d\n", val)
	}

	fmt.Printf("Duration: %v\n", time.Since(start))

	// close the channel when we are done

	go func() {

		for i := 0; i < count; i++ {
			fmt.Printf("Sending second: %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("Received second: %d\n", i)

	}

	// make a buffered channel
	// when buffered, we pass a value indicating how many value we can send without blocking

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}
