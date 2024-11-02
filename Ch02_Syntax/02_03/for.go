package main

import (
	"fmt"
)

func main() {

	count := 5
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	fmt.Println("------------")

	for i := 0; i < count; i++ {

		if i == 3 {
			continue

		}

		fmt.Println(i)
	}

	fmt.Println("------------")

	count = 10

	for {
		if count == 5 {
			break
		}
		fmt.Printf("count = %d\n", count)
		count--
	}

	fmt.Println("------------")
	fmt.Printf("count = %d\n", count)
	for count >= 0 {
		fmt.Printf("count = %d for the second\n", count)
		count--
	}

}
