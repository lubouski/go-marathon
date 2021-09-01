package main

import (
	"fmt"
)

func main() {
	// The comma ok idiom
	m := map[string]int{
		"hello": 5,
		"world": 10,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["human"]
	fmt.Println(v, ok)

	// Deleting from Maps
	delete(m, "hello")

	fmt.Println(m)
}
