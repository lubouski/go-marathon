package main

import (
	"fmt"
)

func main() {
//	prereq := [][]int{{1,0},{2,0},{3,1},{3,2},{3,5}}
	prereq := [][]int{{1,0},{2,0},{3,1},{3,2}}
	numCourses := 4
	weightList := make([]int,numCourses)
	results := []int{}

// Creating weight list for Topological graph search algo

	for _,v := range prereq {
		weightList[v[0]] = weightList[v[0]] + 1
	}

	fmt.Println(weightList)

// Creating graph map from adjacency list

	graph := map[int][]int{}

	for _,v := range prereq {
		graphSlice := graph[v[1]]
		graphSlice = append(graphSlice, v[0])
		graph[v[1]] = graphSlice
	}

// This part is for detecting that graph has no cycles, with a help of depth first search algo

	white := map[int]int{}

	for k,_ := range graph {
		white[k] = k
	}

	gray := map[int]int{}
	black := map[int]int{}

	for len(white) > 0 {
		for _,v := range white {
			if dfs(v, white, gray, black, graph) {
				fmt.Println("Has a loop")
			}
		}
	}
	fmt.Println("No loop")

	loopOver(&weightList, numCourses, prereq, &results)
	fmt.Println(results)
}

func loopOver(weightList *[]int, numCourses int, prereq [][]int, results *[]int) {
        for k,el := range *weightList {
                if el == 0 {
                        *results = append(*results, k)
                        (*weightList)[k] = 99
                        for _,val := range prereq {
                                if val[1] == k {
                                        (*weightList)[val[0]] = (*weightList)[val[0]] - 1
                                }
                        }
			loopOver(weightList, numCourses, prereq, results)
                }
        }
}

func dfs(v int, white map[int]int, gray map[int]int, black map[int]int, graph map[int][]int) bool {
	moveV(v, white, gray)
	for _,neighbor := range graph[v] {
		if _,ok := black[neighbor]; ok {
			continue
		} else if _,ok := gray[neighbor]; ok {
			return true
		} else if dfs(neighbor,white,gray,black,graph) {
			return true
		}
	}
	moveV(v, gray, black)
	return false
}

func moveV(v int, white map[int]int, gray map[int]int) {
	delete(white, v)
	gray[v] = v
}
