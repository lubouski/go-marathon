package main

import "fmt"

type TreeNode struct {
    Left  *TreeNode
    Right *TreeNode
    Value int
}

var DefaultValue int = -1024

func main() {
	tree := &TreeNode{Value: 200, Right: nil, Left: nil}
	Insert(tree, 100)
	Insert(tree, 500)
	Insert(tree, 133)
	Insert(tree, 78)
	fmt.Println(tree)

	fmt.Println(Search(tree, 500))

	Print(tree)
}

func Insert(tree *TreeNode, k int) {
	if tree.Value > k {
		if tree.Right == nil {
			tree.Right = &TreeNode{Value: k}
		}
		Insert(tree.Right, k)
	} else if tree.Value < k {
                if tree.Left == nil {
                        tree.Left = &TreeNode{Value: k}
                }
                Insert(tree.Left, k)
	}
}

func Search(tree *TreeNode, k int) bool {
	if tree.Value > k {
		if tree.Left == nil {
			return false
		}
		 return Search(tree.Left, k)
	} else if tree.Value < k {
		if tree.Right == nil {
			return false
		}
		return Search(tree.Right, k)
	}
	return true
}

func Print(tree *TreeNode) {
	if tree == nil {
		return
	}

	fmt.Println(tree.Value)
	Print(tree.Left)
	Print(tree.Right)
}
