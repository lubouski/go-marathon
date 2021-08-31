package main

import (
	"fmt"
)

func main() {
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))

	// Slices with overlapping storage
	s := []int{1, 2, 3, 4}
	y := s[:2]
	z := s[1:]
	s[1] = 20
	y[0] = 10
	z[1] = 30
	fmt.Println("s:", s)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
