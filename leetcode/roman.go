package main

import (
	"fmt"
)

func main() {
	romanSymb := map[string]int{"I":1,"V":5,"X":10,"L":50,"C":100,"D":500,"M":1000}
	var sum int
	for k,v := range s {
		if k < len(s) - 1 && romanSymb[string(v)] < romanSymb[string(s[k+1])] {
			sum -= romanSymb[string(v)]
		} else {
			sum += romanSymb[string(v)]
		}
	}
	fmt.Println(sum)
}
