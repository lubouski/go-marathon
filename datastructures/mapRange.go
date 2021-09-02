package main

import (
	"fmt"
	"sort"
)

func main() {
	grades := map[string]float32{"Alice": 77.2, "Jon": 82.3, "Tracy": 68.9}
	var names []string
	for name := range grades {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s has a grade of %0.1f%%\n", name, grades[name])
	}
}
