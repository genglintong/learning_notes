package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    middle := split(head)
    
    return merge(sortList(head), sortList(middle))
}

// 切分list
func split(head *ListNode) *ListNode {
    slow, fast := head, head
    var tail *ListNode
    
    for fast != nil && fast.Next != nil {
        tail = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    tail.Next = nil
    return slow
}

func merge(left ,right *ListNode) *ListNode {
    var head, cur, pre *ListNode
    for left != nil && right != nil {
        if left.Val < right.Val {
            cur = left
            left = left.Next
        } else {
            cur = right
            right = right.Next
        }
        // 生成head
        if head == nil {
            head = cur
        } else {
            pre.Next = cur
        }
        pre = cur
    }
    
    if left == nil {
        pre.Next = right
    } else {
        pre.Next = left
    }
    
    return head
}
