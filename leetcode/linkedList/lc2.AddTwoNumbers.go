package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    sum := 0
    carry := 0
    result := &ListNode{}
    head := result
    for l1 != nil || l2 != nil {
        sum = 0
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        sum += carry
        carry = sum / 10
        sum %= 10
        result.Next = &ListNode{Val: sum}
        result = result.Next
    }

    if carry != 0 {
        result.Next = &ListNode{Val: carry}
    }

    return head.Next
}

func main() {
	node := &ListNode{Val: 100}
	node2 := &ListNode{Val: 50}
	fmt.Println(addTwoNumbers(node, node2))
}
