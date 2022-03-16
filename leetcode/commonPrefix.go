package main

import (
	"fmt"
)

func main() {
	//strs := []string{"flower","flow","flight"}
	strs := []string{"a"}

	min := len(strs[0])
	for _,v := range strs {
		if min > len(v) {
			min = len(v)
		}
	}
	var ref string
	var commonPref string
	var commonLetter string
	for i := 0; i < min; i++ {
		ref = string(strs[0][i])
		fmt.Println("ref",ref)
		for _,w := range strs {
			if string(w[i]) == ref {
				commonLetter = string(w[i])
			} else {
				fmt.Println("return ",commonPref)
			}
		}
		commonPref = commonPref + commonLetter
	}
	fmt.Println(commonPref)
}
