package main

import (
	"fmt"
	"strings"
)

func main() {

	book := "The color of magic"

	fmt.Println(len(book))
	fmt.Println(book[0])
	fmt.Println(book[0:4])
	fmt.Println(book[0:])
	fmt.Println(book[:4])
	fmt.Println(book[:])
	//fmt.Println(book[0:5:5])
	fmt.Println(strings.Split(book, ""))
	fmt.Println("Spliting is " + book[0:5])

	// multi line string
	fmt.Println(`
		Hello
		World
	`)
}
