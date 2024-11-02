package main

import "fmt"

func main() {

	values := []int{1, 2, 3, 4, 5}
	doubleAt(values, 2)
	fmt.Println(values)

	value := 10
	double(value)
	fmt.Println(value)
	// this does not changes the input value, even though the slice in the first
	// example changes. the first is a pointer, and it can change

	// to solve this, we can use pointers

	doublePtr(&value)
	fmt.Println(value)

}

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(value int) {
	value *= 2
}

// double that receives a pointer
func doublePtr(value *int) {
	current := *value

	fmt.Print("Current value: ")
	fmt.Println(current)
	*value *= 2
}
