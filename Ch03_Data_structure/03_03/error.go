package main

import (
	"fmt"
	"math"
)

func main() {

	s1, err := sqrt(2.0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s1)
	}

	s2, err := sqrt(-2.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s2)
	}
	fmt.Println("Hello World")
}

func sqrt(x float64) (float64, error) { // return float64 and error
	if x < 0 {
		return 0, fmt.Errorf("norgate math: square root of negative number: %v", x)
	}
	return math.Sqrt(x), nil
}
