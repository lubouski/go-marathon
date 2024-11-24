package main

import (
	"fmt"
)

func findMin(nums []int) int {
	// [4,5,6,7,0,1,2,3]
	// if 7 > 3 we only interested in right part of the slice
	// if 0 < 3 it means high is the lowest
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else if nums[mid] < nums[high] {
			high = mid
		}
	}
	return nums[low]
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMin(nums))
}
