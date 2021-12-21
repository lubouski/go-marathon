package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPowerOfTwo(22))
}

func isPowerOfTwo(n int) bool {
    if n == 1 {
        return true
    } else if n == 0 {
        return false
    } else if n == 2 {
        return true
    } else if n % 2 == 0 {
        count := n / 2
        for count > 0 {
            if count / 2 == 0 {
                return true
            } else if count % 2 != 0 {
                return false
            } else {
                count = count / 2
            }
        }
    } 
    return false
}
