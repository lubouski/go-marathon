package main

import (
	"fmt"
)

func dailyTempretures(tempretures []int) []int {
	result := make([]int, len(tempretures))
	stack := make([][]int,0)
	for idx, temp := range tempretures {
		if len(stack) == 0 {
			stack = append(stack, []int{temp,idx})
			continue
		}
		for stack[len(stack)-1][0] < temp {
			result[stack[len(stack)-1][1]] = idx - stack[len(stack)-1][1]
			stack = stack[:len(stack)-1] // pop
			if len(stack) == 0 {
				break
			}
		}
		stack = append(stack, []int{temp,idx})
	}
	return result
}

func main() {
	tempretures := []int{72,73,69,70,71,74,73,72}

	daily := dailyTempretures(tempretures)

	fmt.Println(daily)
}
