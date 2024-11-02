package main

import (
	"fmt"
)

func main() {

	x := 10

	if x > 5 {
		fmt.Println("X is greater than 5")
	} else {
		fmt.Println("X is less than or equal to 5")
	}

	if x > 100 {
		fmt.Println("X is greater than 100")
	} else {
		fmt.Println("X is less than or equal to 100")
	}

	if x > 5 && x > 30 {
		fmt.Println("X is between 5 and 30")
	}

	if x > 5 || x > 30 {
		fmt.Println("X is greater than 5 or 30")
	}

	switch x {
	case 1:
		fmt.Println("X is 1")

	case 2:
		fmt.Println("X is 2")

	case 3:
		fmt.Println("X is 3")

	default:
		fmt.Println("X is not 1, 2, or 3")
	}

	switch {
	case x > 5:
		fmt.Println("X is greater than 5")
	case x < 5:
		fmt.Println("X is less than 5")
	default:
		fmt.Println("X is not greater than 5")
	}
}
