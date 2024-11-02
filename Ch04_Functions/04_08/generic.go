package main

import (
	"fmt"
)

// generic just like interface, can be a set of tasks
type Ordered interface {
	int | float64 | string
}

func min[T Ordered](values []T) (T, error) {
	
	if len(values) == 0 {
		var zero T
		return zero, fmt.Errorf("cannot get the min of an empty slice")
	}
	min := values[0]
	for _, value := range values[1:] {
		if value < min {
			min = value
		}
	}
	return min, nil
}

func main() {
	
	fmt.Println(min([]int{1, 2, 3, 4, 5}))
	fmt.Println(min([]float64{1.1, 2.2, 3.3, 4.4, 5.5}))
	fmt.Println(min([]string{"a", "b", "c", "d", "e"}))
}