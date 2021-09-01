package main

import (
	"fmt"
)

func hourglassSum(arr [6][6]int32) int32 {
    var sum int32 = -2147483648
    for n := 0; n < 4; n++ {
        for m := 0; m < 4; m++ {
            goldenGlass := arr[n][m] + arr[n][m + 1] + arr[n][m + 2] + arr[n + 1][m + 1] + arr[n + 2][m] + arr[n + 2][m + 1] + arr[n + 2][m + 2]
            if goldenGlass > sum {
                sum = goldenGlass
            }
        }
    }
    return sum
}

func main() {
	var arr[6][6]int32 
	arr[2][2] = 100
	arr[0][0] = 19
	arr[0][1] = 9
	fmt.Println(hourglassSum(arr))
}

