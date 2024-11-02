package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
	Age  int
}

type Budget struct {
	CompainId string
	Balance   float64
	Expires   time.Time
}

func main() {
	p := Person{
		Name: "John Doe",
		Age:  30,
	}
	fmt.Println(p)

	b := Budget{
		"ABC123",
		1000.00,
		time.Now().Add(time.Hour * 24 * 30),
	}
	fmt.Println(b)

	//fmt.Printf("b is of type %T and %#v\n", b, b)

	// create badget by name

	b2 := Budget{
		Balance: 2000.00,
		Expires: time.Now().Add(time.Hour * 24 * 30),
	}
	fmt.Println(b2)

	var b3 Budget
	fmt.Println(b3)

	b3.Balance = 3000.00
	fmt.Println(b3)

	// every attribute that starts with upper case can be accessed outside the
	// package, while attributes that starts with lower case can only be accessed
	// inside the package. In go, we call thesw attribute exported and unexported
}
