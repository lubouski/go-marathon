package main

import (
	"fmt"
)

func findMedian(arr []int32) int32 {
    newArr := make([]int32, len(arr))
    deleteIndex := 0
    for index, _ := range newArr {
        newArr[index] = arr[0]
        for k, _ := range arr {
            if arr[k] <= newArr[index] {
                newArr[index] = arr[k]
                deleteIndex = k
                continue
            }
        }
	// remove minimal index from slice
        arr[deleteIndex] = arr[len(arr)-1]
        arr = arr[:len(arr)-1]
	// assign last element to the newArr
        if len(arr) == 1 {
            newArr[len(newArr)-1] = arr[0]
            break
        }
    }
    return newArr[len(newArr)/2]
}

func main() {
	arr := []int32{0, 1, 2, 4, 6, 5, 3}
	fmt.Println(findMedian(arr))
}


