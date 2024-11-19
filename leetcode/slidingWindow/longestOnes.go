package main

import (
	"fmt"
)

func longestOnes(nums []int, k int) int {
	count := 0
	l := 0
	numsZero := make(map[int]int, 1)
	// while numZero[0] > k
	for r := 0; r < len(nums); r++ {
		if nums[r] == 0 {
			numsZero[0]++
		}
		for {
			if numsZero[0] > k {
				if nums[l] == 0 {
					numsZero[0]--
				}
				l++
				continue
			}
			break
		}
		count = Max(count, (r-l)+1)
	}
	return count
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ones := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	fmt.Println(longestOnes(ones, 2))
}
