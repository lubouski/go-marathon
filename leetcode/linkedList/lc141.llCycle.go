package main

import (
	"fmt"
)

type ListNode struct {
     Val int
     Next *ListNode
}

func hasCycle(head *ListNode) bool {
    nodeTracker := make(map[*ListNode]bool)
    current := head
    for current != nil {
        if nodeTracker[current] {
            return true
        }
        nodeTracker[current] = true
        current = current.Next
    }
    return false
}

func main() {
	node := &ListNode{Val: 100}
	fmt.Println(hasCycle(node))
}
