package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func kSmallest(root * TreeNode, k int) int {
	count := 0
	result := 0

	var inOrder func(node *TreeNode)

        inOrder = func(node *TreeNode) {
		if node == nil || count >= k {
			return
		}
	
		inOrder(node.Left)

		count++
		if count == k {
			result = node.Val
			return
		}

		inOrder(node.Right)
	}
	inOrder(root)
	return result
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 4}

	fmt.Println("Smallest:", kSmallest(root, 2))
}
