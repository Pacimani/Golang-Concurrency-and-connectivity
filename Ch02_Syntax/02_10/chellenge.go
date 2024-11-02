package main

import (
	"fmt"
	"strings"
)

func main() {

	text := `
		Needles and pins
		Needles and pins
		Just need a few
		to c
		To catc me the wind
		`
	words := strings.Fields(text)
	fmt.Println(words)

	counts := map[string]int{} // word -> count

	for _, word := range words {
		counts[strings.ToLower(word)]++
	}

	fmt.Println(counts)
}
