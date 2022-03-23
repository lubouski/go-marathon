package main

import (
	"fmt"
)

func main() {
	stocks := []int{7,1,5,6,2,3}
	left := stocks[0]
	var right, topSell int
	for i := 1; i < len(stocks); i++ {
		right = stocks[i]
		if left > right {
			left = right
		}
		if topSell < right - left {
			topSell = right -left
		}
	}
	fmt.Println(topSell)
}
