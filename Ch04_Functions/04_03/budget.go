package main

import (
	"fmt"
	"time"
)

type Budget struct {
	CompaignId string
	Balance    float64
	Expires    time.Time
}

func NewBudget(compaignId string, balance float64, expires time.Time) (*Budget, error) {

	// data validation
	if balance <= 0 {
		return nil, fmt.Errorf("balance cannot be negative")
	}

	if compaignId == "" {
		return nil, fmt.Errorf("compain id cannot be empty")
	}
	return &Budget{
		CompaignId: compaignId,
		Balance:    balance,
		Expires:    expires,
	}, nil
}

func main() {

	expires := time.Now().Add(time.Hour * 24 * 30)
	b1, err := NewBudget("ABC123", 1000.00, expires)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b1)

	expires = time.Now().Add(time.Hour * 24 * 30)
	b2, err := NewBudget("", 1000.00, expires)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b2)
}
