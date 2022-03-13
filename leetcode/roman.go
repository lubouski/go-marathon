package main

import (
	"fmt"
)

func main() {
	w := "XIV"

	var sum int

	for k,_ := range w {
		if k == len(w) - 2 && str(v) == "I" && (w[k+1] == "V" || w[k+1] == "X") {
			if w[k+1] == "V" {
				sum += 4
				break
			} else {
				sum += 9
				break
			}
		}
	}
	fmt.Println("len", len(w))
}
