package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var maxDiameter int

func diameter(root *TreeNode) int {
	maxDiameter = 0
	height(root)
	return maxDiameter
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := height(node.Left)
	rightHeight := height(node.Right)

	maxDiameter = maxNum(maxDiameter, leftHeight+rightHeight)

	return 1 + maxNum(leftHeight, rightHeight)
}

func maxNum(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 4}

	fmt.Println("Diameter of the tree: ", diameter(root))
}
