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
### Static Types
Define a type like this:
```
type item struct {
  id    int
  value string
}
```
Then define a method for my type:
```
func (i *item) String() string {
  return i.value
}
```
Then define a function that uses my type:
```
func ConcatItems(a item, d item) {
  return a.String() + b.String()
}
```
Method to create new instances of our type:
```
func New(id int, value string) item {  
  i := item{id, value}
  return i
}
```
### JSON
Define how a type should be represented in JSON, then read a JSON string and convert it directly to my type. Validation is automatic.
```
type Mission struct {
  Name     string            `json:"name"`
  Services []Service         `json:"services"`
  Stages   []Stage           `json:"stages"`
  Params   map[string]string `json:"params"`
  isValid  bool
}
func NewFromJSON(jsonString []byte) Mission {
  var m Mission
  json.Unmarshal(jsonString, &m)
  return m
}
```

