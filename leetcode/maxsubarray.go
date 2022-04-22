package main

import (
	"fmt"
)

func main() {
	arr := []int{-1}
	fmt.Println(maxSubArrayV2(arr))
}

func maxSubArrayV2(arr []int) int {
	currentSum := arr[0]
	maxSum := currentSum
	for i := 1; i < len(arr); i++ {
		currentSum = max(arr[i], currentSum + arr[i])
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}


func max(val1 int, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}
