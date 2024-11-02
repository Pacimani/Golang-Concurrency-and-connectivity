package main

import (
	"fmt"
	"time"
)

type Budget struct {
	CompainId string
	Balance   float64
	Expires   time.Time
}

func (b Budget) TimeLeft() time.Duration {
	return b.Expires.Sub(time.Now().UTC())
}

// we pass the pointer to the budge because we want to change the balance inside the badget

func (b *Budget) Update(amount float64) {
	b.Balance += amount
}
func main() {

	b := Budget{
		"ABC123",
		1000.00,
		time.Now().Add(time.Hour * 24 * 30),
	}

	fmt.Println(b.TimeLeft())

	b.Update(2000.00)
	fmt.Println(b)

}
