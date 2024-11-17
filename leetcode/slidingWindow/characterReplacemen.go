package main

import (
	"fmt"
)

func characterReplacement(s string, k int) int {
	count := 0
	l := 0
	letters := make([]int, 26)

	for r, num := range s {
		// update slice with letter
		letters[num - 65]++
		// while w - max(letters) > k
		for {
			if (r - l) + 1 - MaxSl(letters) > k {
				// move left, and remove left letter from slice
				letters[rune(s[l])-65]--
				l++
			}
			break
		}
		// count max letters in window w = (r - l) + 1
		count = Max(count, (r - l) + 1)
	}
	return count
}

func main() {
	s := "AABBAA"
	fmt.Println(characterReplacement(s, 1))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxSl(sl []int) int {
	maximim := 0
	for _, v := range sl {
		if v > maximim {
			maximim = v
		}
	}
	return maximim
}
