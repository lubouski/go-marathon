package main

import (
	"fmt"
)

func rotateLeft(d int32, arr []int32) []int32 {
    for i := int32(0); i < d; i++ {
        indexZero := arr[0]
        for index, _ := range arr {
            if index == len(arr) - 1 {
                arr[index] = indexZero
            } else {
                arr[index] = arr[index + 1]
            }
        }
    }
    return arr
}


func main() {
	var d int32 = 2
	var arr = []int32{1, 2, 3, 4, 5}
	fmt.Println(rotateLeft(d, arr))
}
