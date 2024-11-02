package main

import (
	"fmt"
	"log"
)

// Square struct
type Square struct {
	X, Y, Length int
}

func NewSquare(x, y, length int) (*Square, error) {

	if x <= 0 || y <= 0 || length <= 0 {
		return nil, fmt.Errorf("invalid dimensions")
	}
	return &Square{x, y, length}, nil
}

func (s *Square) Move(dx int, dy int) {
	s.X, s.Y = s.Y + dx, s.Y+dy
}

func (s *Square) Area() int {
	return s.Length * s.Length
}

func main() {

	s, err := NewSquare(1, 1, 10)
	if err != nil {
		log.Fatal(err)
	}

	s.Move(2, 3)
	
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())
}