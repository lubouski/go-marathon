package main

import (
	"fmt"
)

func searchMatrix(matrix [3][4]int, target int) bool {
	r, k := len(matrix), len(matrix[0])
	// flatten the matrix to 12 elements and 11 indexes
	low, high := 0, r*k-1
	for low <= high {
		mid := low + (high-low)/2
		n, m := findMNIndex(mid, r, k)
		val := matrix[n][m]
		if val < target {
			low = mid + 1
		} else if val > target {
			high = mid - 1
		} else {
			return true
		}
	}
	return false
}

func findMNIndex(number, n, m int) (r, k int) {
	// r - rows, k - columns
	r = (number / m)
	k = (number % m)
	return r, k
}

func main() {
	matrix := [3][4]int{
		{1, 2, 3, 4},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}

	fmt.Println(searchMatrix(matrix, 7))
}
