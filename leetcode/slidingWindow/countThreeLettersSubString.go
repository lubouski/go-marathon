package main

import (
	"fmt"
)

func countGoodSubstrings(s string) int {
	l := 0
	count := 0
	r := 2

	for r < len(s) {
		first_char := s[l]
		second_char := s[l+1]
		third_char := s[r]
		if first_char != second_char && first_char != third_char && second_char != third_char {
			count++
		}
		l++
		r++
	}
	return count
}

func main() {
	s := "xyzzaz"
	fmt.Println(countGoodSubstrings(s))
}
