package main

import (
	"fmt"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    slowPointer := dummy
    fastPointer := dummy
    for i := 0; i <= n; i++ {
        fastPointer = fastPointer.Next
    }
    for fastPointer != nil {
        slowPointer = slowPointer.Next
        fastPointer = fastPointer.Next
    }
    slowPointer.Next = slowPointer.Next.Next
    return dummy.Next
}

func main() {
	node := &ListNode{val: 100}
	fmt.Println(removeNthFromEnd(node, 1))	
}
