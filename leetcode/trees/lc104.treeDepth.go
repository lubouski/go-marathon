package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	right := maxDepth(root.Right)
	left := maxDepth(root.Left)

	return 1 + maxNum(right, left)
}

func maxNum(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 2}

	fmt.Println(maxDepth(root))
}
