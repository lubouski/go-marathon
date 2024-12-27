package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, balanced := checkBalance(root)
	return balanced
}

func checkBalance(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	leftHeight, leftBalanced := checkBalance(node.Left)
	if !leftBalanced {
		return 0, false
	}

	rightHeight, rightBalanced := checkBalance(node.Right)
	if !rightBalanced {
		return 0, false
	}

	if absDiffInt(leftHeight, rightHeight) > 1 {
		return 0, false
	}

	return 1 + maxH(leftHeight, rightHeight), true
}

func maxH(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Left.Left = &TreeNode{Val: 5}

	fmt.Println(isBalanced(root))
}
