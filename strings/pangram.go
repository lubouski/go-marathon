package main

import (
	"fmt"
	"strings"
)

func pangrams(s string) string {
    s = strings.ReplaceAll(s, " ", "")
    s = strings.ToLower(s)
    m := map[string]string{}
    for _, j := range s {
        m[string(j)] = string(j)
    }
    if len(m) == 26 {
        return "pangram"
    } else {
        return "not pangram"
    }
}

func main() {
	s := "We promptly judged antique ivory buckles for the next prize"
	fmt.Println(pangrams(s))
}
