package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode // Previous node starts as nil
	curr := head       // Current node starts at the head

	// Traverse the list
	for curr != nil {
		next := curr.Next // Save the next node

		curr.Next = prev // Reverse the link

		// Move pointers forward
		prev = curr // Move prev to the current node
		curr = next // Move curr to the next node
	}

	// prev is now the new head of the reversed list
	return prev
}

func main() {
	node := &ListNode{val: 100}
	fmt.Println(reverseList(node))
}
