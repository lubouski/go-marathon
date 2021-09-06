package main

import (
	"fmt"
	"strings"
)

func caesarCipher(s string, k int32) string {
    var sl []string
    kk := rune(k % int32(26))
    for _, j := range s {
        if j < 65 || j > 90 && j < 97 || j > 122 {
            sl = append(sl, string(j))
        } else if j < 91 && j+kk > 90 {
            sl = append(sl, string(j-26+kk))
        } else if j < 91  {
            sl = append(sl, string(j+kk))
        } else if j+kk > 122 {
            sl = append(sl, string(j-26+kk))
        } else {
            sl = append(sl, string(j+kk))
        }
    }
    return strings.Join(sl, "")
}

func main() {
	s := "AbcdeF"
	var k int32 = 2
	fmt.Println(caesarCipher(s, k))
}
