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
    fast := head
    slow := head

    for fast != nil && slow != nil {
        fast = fast.Next
        slow = slow.Next

        if fast == nil || slow == nil {
            return false
        }
        slow = slow.Next

        if slow == fast {
            return true
        }
    }
    return false
}
