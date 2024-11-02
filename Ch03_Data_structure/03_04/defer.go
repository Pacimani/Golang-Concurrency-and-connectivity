package main

import "fmt"

func main() {
	r1, err := acquire("foo")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer release(r1)

	r2, err := acquire("bar")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer release(r2)

	fmt.Println("Worker")

}

func acquire(name string) (string, error) {

	return name, nil
}

func release(name string) {
	fmt.Println("Releasing: ", name)
}
