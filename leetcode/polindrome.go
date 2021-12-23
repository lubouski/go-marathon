package main

import (
	"fmt"
)

func main() {
	isPalindrome(2442)
}

func isPalindrome(x int) {
    if x == 0 {
        fmt.Println(true)
    } else if x < 0 {
        fmt.Println(false)
    } else if x > 0 {
        poli := 0
	    equal := x
	    for {
		    poli = poli * 10 + x % 10
		    x = (x - x % 10) / 10
		    if x < 1 {
			    break
		    }
	    }
	    if equal == poli {
		    fmt.Println(true)
	    } else {
		fmt.Println(false)
	    }
    }
}
