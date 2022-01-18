package main

import (
	"fmt"
)

func main() {
	sl := []int{2,1,4,3,7,6,5}
	selectionSort(sl)
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n - 1; i++ {
		minInx := i
		for j := i +1; j < n; j++ {
			if arr[j] < arr[minInx] {
				arr[j], arr[minInx] = arr[minInx], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
