package main

import (
	"fmt"
)

func main() {
	target := 5
	nums := []int{1,3,5,6}
	fmt.Println(searchInsert(nums,target))
}

func searchInsert(nums []int, target int) int {
	var n int
	for len(nums) != 0 {
		if target == nums[len(nums)/2] {
			for k,v := range nums {
				if v == target {
					n = n + k
				}
			}
			break
		} else if target > nums[len(nums)/2] {
			n = n + len(nums)/2
			nums = nums[len(nums)/2:]
			if len(nums) == 1 {
				return n + 1
			}
		} else if target < nums[len(nums)/2] {
			n = n + 0
			nums = nums[:len(nums)/2]
			if len(nums) == 1 {
				if target < nums[0] || target == nums[0] {
					return n
				}
				return n +1
			}
		}
	}
	return n
}
