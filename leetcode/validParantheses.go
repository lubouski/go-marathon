package main

import (
	"fmt"
)

func main() {
	// Parantheses string
	s := "(){}}{"
	val := map[string]string{
		")":"(",
		"]":"[",
		"}":"{",
	}
	// map to check equal quatity for parantheses
	check := map[string]int{}
	// array to convert a string
	arr := make([]string, len(s))
	for k,v := range s {
		check[string(v)] = check[string(v)] + 1
		arr[k] = string(v)
	}
	if check["("] != check[")"] || check["{"] != check["}"] || check["["] != check["]"] {
		fmt.Println("false")
	}
	// edge case if string starts with closed 
	if arr[0] == ")" || arr[0] == "}" || arr[0] == "]" {
		fmt.Println("false 2")
	}
	// stack buffer to track right order for closing brackets
	var buff []string
	for _,j := range arr {
		if j == "(" || j == "[" || j == "{" {
			buff = append(buff,j)
		} else if j == ")" || j == "]" || j == "}" {
			if len(buff) == 0 {
				fmt.Println("false 11")
			} else if buff[len(buff)-1] != val[j] {
				fmt.Println("false 1")
			}
			buff = buff[:len(buff)-1]
		}
	}
	if len(buff) == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
