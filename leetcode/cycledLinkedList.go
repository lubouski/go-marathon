package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    indexMap := map[*ListNode]int{}
    for head != nil {
        if indexMap[head] == 1 {
            return true
        }
        indexMap[head] = 1
        head = head.Next
    }
    return false
}

func main() {
	mylist := ListNode{}
	fmt.Println(hasCycle(mylist))
}
