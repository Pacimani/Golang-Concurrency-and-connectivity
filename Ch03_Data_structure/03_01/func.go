package main

import (
	"fmt"
	"math"
)

func main() {

	val := add(1, 2)
	fmt.Println(math.Sqrt(float64(val)))

	div, mod := divmod(10, 3)
	fmt.Println(div, mod)

}

// add a to b
func add(x int, y int) int {
	return x + y
}

// divmod returns quotient and remainder
func divmod(x int, y int) (int, int) {
	return x / y, x % y
}
