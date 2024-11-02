package main

import (
	"fmt"
	"math"
)

// Square struct
type Square struct {
	Length float64
}

// interface is a very important concept/abstraction
func (s Square) Move(dx int, dy int) {
	s.Length = s.Length + float64(dx) + float64(dy)
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}


func (c Circle) Area() float64 {
	return math.Pi * float64(c.Radius) * float64(c.Radius)
}


// sumAreas returns the sum of all areas in the slice
func sumAreas(s []Shape) float64 {

	sum := 0.0
	for _, shape := range s {
		sum += shape.Area()
	}
	return sum
}


type Shape interface {
	Area() float64
}
func main() {

	s := Square{20}
	fmt.Println(s.Area())

	c := Circle{10}
	fmt.Println(c.Area())

	shapes := []Shape{c, s}
	total := sumAreas(shapes)
	fmt.Println(total)
	
	
}