## Leetcode examples on data structures and algorithms
As an education joarney, leetcode is great resource to strengthen computer scirnce knowledge.

### Breadth-First search
Mine realisation of breadth first search algorithm. It's based on the queue structure and graphs. This particular example requires to have acyclic graph to iterate on. 
```
func main() {
	graph := map[string][]string{
		"Bill": []string{"Jill","Fred","John"},
		"Fred": []string{"Tom"},
		"Jill": []string{"Oscar"},
		"John": []string{"Scott"},
	}

	var person string

	queue := graph["Bill"]
	// while queue is not empty pick first element and check, if not add all peers to the queue
	for len(queue) != 0 {
		person = queue[0]
		if person == "Scott" {
			fmt.Println("Hey Scott")
		}
		queue = queue[1:]
		for _,v := range graph[person] {
			queue = append(queue, v)
		}
		fmt.Println(queue)
	}
}
```
