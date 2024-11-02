package main

import "fmt"

func main() {

	//var x float64 // you can change the type to int, float32 etc
	//var y float64

	// x = 1
	// y = 2

	x, y := 1.2, 1.5

	fmt.Printf("x=%v, type of %T\n", x, x)

	fmt.Printf("y=%v, type of %T\n", y, y)

	mean := (x + y) / 2.0

	fmt.Printf("mean=%v, type of %T\n", mean, mean)
}
