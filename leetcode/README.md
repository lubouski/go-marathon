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
### Recursion
Base case usage, it's where you go down the ladder and have condition for the bottom result i.e base case, and then gradually claimbing the ladder. 
```
func main() {
	fmt.Println(sumNums(10))
}

func sumNums(n int) int {
	if n == 1 {
		return 1
	}
	return sumNums(n - 1) + n
}
```

### BinaryTree preorder traversal
In particular case bellow there are two functions, wrapper function `preorderTraversal` is needed to output slice of traversed values.
```
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    sl := []int{}
    preorder(root, &sl)
    return sl
}

func preorder(t *TreeNode, sl *[]int) {
    if t == nil {
        return
    }
    *sl = append(*sl, t.Val)
    preorder(t.Left, sl)
    preorder(t.Right, sl)
}
```

### Remove element from LinkedList
```
func removeElements(head *ListNode, val int) *ListNode {
    // dummy appended Node to help with two pointer (prev and curr)
    dummy := &ListNode{Next: head}
    // srt two pointers prev and curr
    prev, curr := dummy, head

    // iterating over curr to not mess with head
    for curr != nil {
        nxt := curr.Next
        if curr.Val == val {
            prev.Next = nxt
        } else {
            prev = curr
        }
    // assigning curr to nxt to iterate over LinkedList
        curr = nxt
    }
    return dummy.Next
}
```
