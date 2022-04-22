package main

import (
	"fmt"
)

func main() {
	str := " baby  fly    "
	fmt.Println(lengthOfLastWord(str))
}

func lengthOfLastWord(s string) int {
	var count int
	var sw int
	for i := (len(s) - 1) ; i > -1; i-- {
		if string(s[i]) == " " && sw == 0 {
			sw = 0
			continue
		} else if string(s[i]) == " " && sw == 1 {
			break
		}
		sw = 1
		count++
	}
	return count
}
