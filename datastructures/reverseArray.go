package main

import (
	"fmt"
)

func reverseArray(a []int32) []int32 {
 //   var revArr [len(a)]int32
    revArr := make([]int32, len(a))
    revIndex := int32(0)

    for index := len(a) -1; index >= 0; index-- {
        revArr[revIndex] = a[index]
        revIndex++
    }
    return revArr
}

func main() {
	var a []int32
	a = []int32{1, 2, 3, 4, 5}
	fmt.Println("Array", a)
	fmt.Println("Reverse Array", reverseArray(a))
}
