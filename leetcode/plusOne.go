package main

import (
	"fmt"
)

func main() {
	digits := []int{1,2,9,9}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			return digits
		}
		digits[i] = 0
	}
	newDigits := make([]int, len(digits) + 1, len(digits) + 1)
	newDigits[0] = 1
	return newDigits
}
