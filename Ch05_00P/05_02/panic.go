package main

import "fmt"

func safeValue(values []int, index int) (n int, err error) {

	defer func() {

		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
			//fmt.Println(err)
		}
	}()

	return values[index], nil
}

func main() {

	vals := []int{1, 2, 3}

	// v := vals[4] // this will cause pannic, because index is out of bounds. thre is no value at index 4

	v, err := safeValue(vals, 4)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	// fmt.Println(v)
	fmt.Println("Find the index value ", v)
}
