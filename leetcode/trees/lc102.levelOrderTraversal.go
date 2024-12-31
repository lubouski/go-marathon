package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	var result [][]int

	for len(queue) > 0 {
		levelSize := len(queue)
		var curLevel []int

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			curLevel = append(curLevel, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, curLevel)
	}
	return result
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	fmt.Println(levelOrder(root))
}
