package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	letters := make(map[byte]bool)
	count := 0
	max := 0
	l := 0
	n := len(s)
	for r := 0; r < n; r++ {
                // while s[r] in set
		for {
			if _, ok := letters[s[r]]; ok {
				// remove l
				delete(letters, s[l])
				l++
				continue
			}
			break
		}
		count = (r - l) + 1
		max = Max(max, count)
		// add r if it not in set
		letters[s[r]] = true
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
