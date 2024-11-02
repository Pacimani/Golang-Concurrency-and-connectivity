package main

import "fmt"

func main() {

	// slice is a sequence of items, equivalent to an array or list
	// in other programming languages
	loons := []string{"bugs", "daffy", "taz", "l4", "l5"}

	fmt.Printf("Looms is %v (type %T)\n", loons, loons)

	fmt.Println(len(loons))

	fmt.Println(loons[0])
	fmt.Println(loons[1:])

	// you can use loop
	for i := 1; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	for i, name := range loons {
		fmt.Println(i, name)
	}

	// using append
	loons = append(loons, "l6")
	fmt.Println(loons)

}
