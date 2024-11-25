package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] < nums[high] {
			high = mid
		} else if nums[mid] > nums[high] {
			low = mid + 1
		}
	}
	minIndex := high

	var l, r int

	// cover 3 cases: not rotated array, target in first half, target in second part

	if minIndex == 0 {
		l, r = 0, len(nums)-1
	} else if target >= nums[0] && target <= nums[minIndex-1] {
		l, r = 0, minIndex-1
	} else {
		l, r = minIndex, len(nums)-1
	}

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 8, 50, 0, 1, 2, 3}
	fmt.Println(search(nums, 50))
}
