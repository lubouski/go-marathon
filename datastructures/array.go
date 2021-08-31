package main

import (
	"fmt"
)

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "sol", "la", "ti"}
	fmt.Println("Len of array is:", len(notes))
	fmt.Println("Capability of array is:", cap(notes))

	numbers := [3]int{1, 2, 3}
	numbers2 := [3]int{1, 2, 10}

	if numbers == numbers2 {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	// array could be initialized without a number and used ... instead

	var x = [...]int{1, 2, 3}
	var y = [3]int{1, 2, 3}
	fmt.Println(x == y) // prints true
}
