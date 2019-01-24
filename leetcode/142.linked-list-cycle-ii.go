package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    fast := head
    slow := head

    for fast != nil && slow != nil {
        fast = fast.Next
        slow = slow.Next

        if slow == nil || fast == nil{
            return nil
        }

        slow = slow.Next
        if slow == fast {
            fast = head

            for fast != slow {
                fast = fast.Next
                slow = slow.Next
            }
            return slow
        }
    }
    return nil
}
