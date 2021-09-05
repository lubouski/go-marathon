package main

import (
	"fmt"
	"strings"
)

func camelcase(s string) int32 {
    sum := int32(0)
    for _, i := range s {
        if string(i) == strings.ToUpper(string(i)) {
            sum++
        }
    }
    return sum + int32(1)
}

func main() {
	s := "saveChangesInTheEditor"
	fmt.Println("Words count:", camelcase(s))
}
