package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	n := 42
	s := fmt.Sprintf("%d", n) // this turns 42 into string

	fmt.Printf("%s\n", s)

	count := 0

	for a := 1000; a <= 9999; a++ {

		for b := a; b <= 9999; b++ { // don't count twice
			n := a * b

			// if a * b is ended
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				count++
			}
		}
	}

	fmt.Println(count)
}
