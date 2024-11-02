package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 42
	}()

	go func() {
		ch2 <- 48
	}()

	select {
	case x := <-ch1:
		fmt.Printf("ch1: %d\n", x)
	case y := <-ch2:
		fmt.Printf("ch2: %d\n", y)
	}

	fmt.Println("-------")

	out := make(chan float64)

	go func() {
		time.Sleep(100 * time.Millisecond)
		out <- 3.14
	}()

	select {
	case x := <-out:
		fmt.Printf("got  %f\n", x)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("timeout")
	}
}
